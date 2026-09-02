package compressor

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"cobagolang/backend/internal/model"
	"cobagolang/backend/internal/utils"
)

// QPDFCompressor implements PDFCompressor using QPDF
type QPDFCompressor struct {
	binaryPath string
}

// NewQPDFCompressor creates a new QPDF compressor instance
func NewQPDFCompressor(customBinaryPath string) *QPDFCompressor {
	binary := FindExecutable(
		customBinaryPath,
		"qpdf",
		"qpdf.exe",
	)
	return &QPDFCompressor{
		binaryPath: binary,
	}
}

// Name returns the compressor name
func (q *QPDFCompressor) Name() string {
	return "qpdf"
}

// IsAvailable checks if QPDF executable was resolved
func (q *QPDFCompressor) IsAvailable() bool {
	return q.binaryPath != ""
}

// GetBinaryPath returns the resolved executable path
func (q *QPDFCompressor) GetBinaryPath() string {
	return q.binaryPath
}

// Compress executes QPDF with appropriate stream and flate compression flags
func (q *QPDFCompressor) Compress(ctx context.Context, inputPath, outputPath string, level model.CompressionLevel) error {
	log := utils.GetLogger()

	if !q.IsAvailable() {
		return fmt.Errorf("%w: qpdf binary not found", ErrCompressorNotAvailable)
	}

	args := []string{
		"--linearize",
		"--compress-streams=y",
	}

	switch level {
	case model.LevelStudioMaster, model.LevelHighFidelity, model.LevelLow:
		// Basic stream compression
	case model.LevelRecommended, model.LevelMedium, model.LevelCustomTarget:
		// Recompress flate streams
		args = append(args, "--recompress-flate")
	case model.LevelExtreme, model.LevelHigh, model.LevelUltraExtreme:
		// Recompress flate and generate object streams
		args = append(args, "--recompress-flate", "--object-streams=generate")
	default:
		return fmt.Errorf("%w: %s", ErrInvalidLevel, level)
	}

	args = append(args, inputPath, outputPath)

	log.WithFields(map[string]interface{}{
		"engine": "qpdf",
		"binary": q.binaryPath,
		"level":  level,
		"args":   args,
	}).Debug("Executing QPDF command")

	cmd := exec.CommandContext(ctx, q.binaryPath, args...)

	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		log.WithFields(map[string]interface{}{
			"error":  err.Error(),
			"stderr": stderr.String(),
			"stdout": stdout.String(),
		}).Error("QPDF execution failed")

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

// Merge combines multiple PDF files in sequential order using QPDF
func (q *QPDFCompressor) Merge(ctx context.Context, inputPaths []string, outputPath string) error {
	log := utils.GetLogger()

	if !q.IsAvailable() {
		return fmt.Errorf("%w: qpdf binary not found", ErrCompressorNotAvailable)
	}

	if len(inputPaths) == 0 {
		return fmt.Errorf("no input files provided for merge")
	}

	args := []string{"--empty", "--pages"}
	args = append(args, inputPaths...)
	args = append(args, "--", outputPath)

	log.WithFields(map[string]interface{}{
		"engine":      "qpdf",
		"binary":      q.binaryPath,
		"files_count": len(inputPaths),
		"output":      outputPath,
	}).Debug("Executing QPDF Merge command")

	cmd := exec.CommandContext(ctx, q.binaryPath, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		utils.RemoveFileSafely(outputPath)
		return fmt.Errorf("qpdf merge failed: %w - %s", err, stderr.String())
	}

	info, err := os.Stat(outputPath)
	if err != nil || info.Size() == 0 {
		utils.RemoveFileSafely(outputPath)
		return fmt.Errorf("merged output file is missing or empty")
	}

	return nil
}

// Split extracts specific page ranges or splits into single pages using QPDF
func (q *QPDFCompressor) Split(ctx context.Context, inputPath, outputDir string, splitMode string, pageRanges string) ([]string, error) {
	log := utils.GetLogger()

	if !q.IsAvailable() {
		return nil, fmt.Errorf("%w: qpdf binary not found", ErrCompressorNotAvailable)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create split output dir: %w", err)
	}

	var generatedFiles []string

	if splitMode == "range" && pageRanges != "" {
		outPath := filepath.Join(outputDir, fmt.Sprintf("extracted_%s.pdf", strings.ReplaceAll(pageRanges, ",", "_")))
		args := []string{inputPath, "--pages", inputPath, pageRanges, "--", outPath}

		cmd := exec.CommandContext(ctx, q.binaryPath, args...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("qpdf split range failed: %w - %s", err, stderr.String())
		}

		if info, err := os.Stat(outPath); err == nil && info.Size() > 0 {
			generatedFiles = append(generatedFiles, outPath)
		}
	} else {
		outputPattern := filepath.Join(outputDir, "page-%d.pdf")
		args := []string{"--split-pages", inputPath, outputPattern}

		cmd := exec.CommandContext(ctx, q.binaryPath, args...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("qpdf split-pages failed: %w - %s", err, stderr.String())
		}

		matches, err := filepath.Glob(filepath.Join(outputDir, "page-*.pdf"))
		if err != nil || len(matches) == 0 {
			return nil, fmt.Errorf("no split pages generated by qpdf")
		}
		generatedFiles = matches
	}

	log.WithField("count", len(generatedFiles)).Debug("QPDF split completed")
	return generatedFiles, nil
}
