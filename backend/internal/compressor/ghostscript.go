package compressor

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"cobagolang/backend/internal/model"
	"cobagolang/backend/internal/utils"
)

// GhostscriptCompressor implements PDFCompressor using Ghostscript
type GhostscriptCompressor struct {
	binaryPath string
}

// NewGhostscriptCompressor creates a new Ghostscript compressor instance
func NewGhostscriptCompressor(customBinaryPath string) *GhostscriptCompressor {
	binary := FindExecutable(
		customBinaryPath,
		"gswin64c",
		"gswin32c",
		"gs",
		"gswin64c.exe",
		"gswin32c.exe",
		"gs.exe",
	)
	return &GhostscriptCompressor{
		binaryPath: binary,
	}
}

// Name returns the compressor name
func (g *GhostscriptCompressor) Name() string {
	return "ghostscript"
}

// IsAvailable checks if Ghostscript executable was resolved
func (g *GhostscriptCompressor) IsAvailable() bool {
	return g.binaryPath != ""
}

// GetBinaryPath returns the resolved executable path
func (g *GhostscriptCompressor) GetBinaryPath() string {
	return g.binaryPath
}

// Compress executes Ghostscript with appropriate optimization parameters
func (g *GhostscriptCompressor) Compress(ctx context.Context, inputPath, outputPath string, level model.CompressionLevel) error {
	log := utils.GetLogger()

	if !g.IsAvailable() {
		return fmt.Errorf("%w: ghostscript binary not found", ErrCompressorNotAvailable)
	}

	// Determine PDF settings based on compression level
	var pdfSettings string
	switch level {
	case model.LevelStudioMaster:
		pdfSettings = "-dPDFSETTINGS=/prepress"
	case model.LevelHighFidelity, model.LevelLow:
		pdfSettings = "-dPDFSETTINGS=/printer"
	case model.LevelRecommended, model.LevelMedium:
		pdfSettings = "-dPDFSETTINGS=/ebook"
	case model.LevelExtreme, model.LevelHigh:
		pdfSettings = "-dPDFSETTINGS=/screen"
	case model.LevelUltraExtreme:
		pdfSettings = "-dPDFSETTINGS=/screen"
	case model.LevelCustomTarget:
		pdfSettings = "-dPDFSETTINGS=/ebook"
	default:
		return fmt.Errorf("%w: %s", ErrInvalidLevel, level)
	}

	// Base optimization arguments
	args := []string{
		"-sDEVICE=pdfwrite",
		"-dCompatibilityLevel=1.4",
		pdfSettings,
		"-dNOPAUSE",
		"-dQUIET",
		"-dBATCH",
		"-dDetectDuplicateImages=true",
		"-dCompressFonts=true",
		"-dEmbedAllFonts=true",
		"-dSubsetFonts=true",
		"-dAutoRotatePages=/None",
		"-dColorImageDownsampleType=/Bicubic",
		"-dGrayImageDownsampleType=/Bicubic",
		"-dMonoImageDownsampleType=/Bicubic",
		"-dDownsampleColorImages=true",
		"-dDownsampleGrayImages=true",
		"-dDownsampleMonoImages=true",
	}

	switch level {
	case model.LevelStudioMaster:
		args = append(args,
			"-dColorImageResolution=300",
			"-dGrayImageResolution=300",
			"-dMonoImageResolution=300",
		)
	case model.LevelHighFidelity, model.LevelLow:
		args = append(args,
			"-dColorImageResolution=220",
			"-dGrayImageResolution=220",
			"-dMonoImageResolution=220",
		)
	case model.LevelRecommended, model.LevelMedium:
		args = append(args,
			"-dColorImageResolution=150",
			"-dGrayImageResolution=150",
			"-dMonoImageResolution=150",
			"-dOptimize=true",
		)
	case model.LevelExtreme, model.LevelHigh:
		args = append(args,
			"-dColorImageResolution=72",
			"-dGrayImageResolution=72",
			"-dMonoImageResolution=72",
			"-dOptimize=true",
		)
	case model.LevelUltraExtreme:
		args = append(args,
			"-dColorImageResolution=50",
			"-dGrayImageResolution=50",
			"-dMonoImageResolution=50",
			"-dOptimize=true",
		)
	case model.LevelCustomTarget:
		args = append(args,
			"-dColorImageResolution=100",
			"-dGrayImageResolution=100",
			"-dMonoImageResolution=100",
			"-dOptimize=true",
		)
	}

	args = append(args, fmt.Sprintf("-sOutputFile=%s", outputPath), inputPath)

	log.WithFields(map[string]interface{}{
		"engine":      "ghostscript",
		"binary":      g.binaryPath,
		"level":       level,
		"pdfSettings": pdfSettings,
	}).Debug("Executing Ghostscript command")

	cmd := exec.CommandContext(ctx, g.binaryPath, args...)

	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		log.WithFields(map[string]interface{}{
			"error":  err.Error(),
			"stderr": stderr.String(),
			"stdout": stdout.String(),
		}).Error("Ghostscript execution failed")

		// Remove corrupted partial output file if generated
		utils.RemoveFileSafely(outputPath)
		return fmt.Errorf("%w: %s", ErrCompressionFailed, stderr.String())
	}

	// Verify output file exists and is non-empty
	info, err := os.Stat(outputPath)
	if err != nil || info.Size() == 0 {
		utils.RemoveFileSafely(outputPath)
		return fmt.Errorf("%w: output file was not created or is empty", ErrCompressionFailed)
	}

	return nil
}

// Merge combines multiple PDF files in sequential order
func (g *GhostscriptCompressor) Merge(ctx context.Context, inputPaths []string, outputPath string) error {
	log := utils.GetLogger()

	if !g.IsAvailable() {
		return fmt.Errorf("%w: ghostscript binary not found", ErrCompressorNotAvailable)
	}

	if len(inputPaths) == 0 {
		return fmt.Errorf("no input files provided for merge")
	}

	args := []string{
		"-sDEVICE=pdfwrite",
		"-dCompatibilityLevel=1.4",
		"-dNOPAUSE",
		"-dQUIET",
		"-dBATCH",
		fmt.Sprintf("-sOutputFile=%s", outputPath),
	}
	args = append(args, inputPaths...)

	log.WithFields(map[string]interface{}{
		"engine":      "ghostscript",
		"binary":      g.binaryPath,
		"files_count": len(inputPaths),
		"output":      outputPath,
	}).Debug("Executing Ghostscript Merge command")

	cmd := exec.CommandContext(ctx, g.binaryPath, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		utils.RemoveFileSafely(outputPath)
		return fmt.Errorf("merge failed: %w - %s", err, stderr.String())
	}

	info, err := os.Stat(outputPath)
	if err != nil || info.Size() == 0 {
		utils.RemoveFileSafely(outputPath)
		return fmt.Errorf("merged output file is missing or empty")
	}

	return nil
}

// Split extracts specific page ranges or splits into single pages
func (g *GhostscriptCompressor) Split(ctx context.Context, inputPath, outputDir string, splitMode string, pageRanges string) ([]string, error) {
	log := utils.GetLogger()

	if !g.IsAvailable() {
		return nil, fmt.Errorf("%w: ghostscript binary not found", ErrCompressorNotAvailable)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create split output dir: %w", err)
	}

	var generatedFiles []string

	if splitMode == "range" && pageRanges != "" {
		// Example: "1-3" or "2"
		parts := strings.Split(strings.TrimSpace(pageRanges), "-")
		startPage := 1
		endPage := 1
		if len(parts) >= 1 {
			if s, err := strconv.Atoi(strings.TrimSpace(parts[0])); err == nil && s > 0 {
				startPage = s
				endPage = s
			}
		}
		if len(parts) >= 2 {
			if e, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil && e >= startPage {
				endPage = e
			}
		}

		outFileName := fmt.Sprintf("extracted_pages_%d_%d.pdf", startPage, endPage)
		if startPage == endPage {
			outFileName = fmt.Sprintf("extracted_page_%d.pdf", startPage)
		}
		outputPath := filepath.Join(outputDir, outFileName)

		args := []string{
			"-sDEVICE=pdfwrite",
			"-dCompatibilityLevel=1.4",
			"-dNOPAUSE",
			"-dQUIET",
			"-dBATCH",
			fmt.Sprintf("-dFirstPage=%d", startPage),
			fmt.Sprintf("-dLastPage=%d", endPage),
			fmt.Sprintf("-sOutputFile=%s", outputPath),
			inputPath,
		}

		cmd := exec.CommandContext(ctx, g.binaryPath, args...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("split range failed: %w - %s", err, stderr.String())
		}

		if info, err := os.Stat(outputPath); err == nil && info.Size() > 0 {
			generatedFiles = append(generatedFiles, outputPath)
		}
	} else {
		// Split all pages into individual files
		outputPattern := filepath.Join(outputDir, "page_%03d.pdf")
		args := []string{
			"-sDEVICE=pdfwrite",
			"-dCompatibilityLevel=1.4",
			"-dNOPAUSE",
			"-dQUIET",
			"-dBATCH",
			fmt.Sprintf("-sOutputFile=%s", outputPattern),
			inputPath,
		}

		cmd := exec.CommandContext(ctx, g.binaryPath, args...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("split all pages failed: %w - %s", err, stderr.String())
		}

		// Collect generated files
		matches, err := filepath.Glob(filepath.Join(outputDir, "page_*.pdf"))
		if err != nil || len(matches) == 0 {
			return nil, fmt.Errorf("no split pages were generated")
		}
		generatedFiles = matches
	}

	log.WithField("count", len(generatedFiles)).Debug("Ghostscript split completed")
	return generatedFiles, nil
}
