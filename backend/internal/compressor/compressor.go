package compressor

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"cobagolang/backend/internal/model"
)

var (
	ErrCompressorNotAvailable = errors.New("selected PDF compressor engine is not installed or not in PATH")
	ErrCompressionFailed      = errors.New("PDF compression process failed")
	ErrInvalidLevel           = errors.New("invalid compression level specified")
)

// PDFCompressor defines the contract for PDF compression, merging, and splitting engines
type PDFCompressor interface {
	// Compress executes PDF compression on input file and writes to output file
	Compress(ctx context.Context, inputPath, outputPath string, level model.CompressionLevel) error
	// Merge combines multiple PDF files in order into a single output file
	Merge(ctx context.Context, inputPaths []string, outputPath string) error
	// Split extracts specific page ranges or splits into single pages
	Split(ctx context.Context, inputPath, outputDir string, splitMode string, pageRanges string) ([]string, error)
	// Name returns the identifier of the engine
	Name() string
	// IsAvailable checks if the underlying engine binary is accessible
	IsAvailable() bool
	// GetBinaryPath returns the resolved executable path
	GetBinaryPath() string
}

// FindExecutable searches for the first available binary name from a list of candidate names
func FindExecutable(configuredPath string, candidates ...string) string {
	if configuredPath != "" {
		if path, err := exec.LookPath(configuredPath); err == nil {
			return path
		}
		if _, err := os.Stat(configuredPath); err == nil {
			return configuredPath
		}
		return configuredPath
	}

	for _, name := range candidates {
		if path, err := exec.LookPath(name); err == nil {
			return path
		}
	}

	// Search common Windows directories
	commonDirs := []string{
		`C:\Program Files\gs`,
		`C:\Program Files (x86)\gs`,
		`C:\Program Files\qpdf`,
		`C:\Program Files (x86)\qpdf`,
	}

	for _, baseDir := range commonDirs {
		var found string
		_ = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if !info.IsDir() {
				for _, cand := range candidates {
					if strings.EqualFold(filepath.Base(path), cand) || strings.EqualFold(filepath.Base(path), cand+".exe") {
						found = path
						return filepath.SkipAll
					}
				}
			}
			return nil
		})
		if found != "" {
			return found
		}
	}

	return ""
}

// NewCompressor initializes the appropriate PDF compressor based on configuration
func NewCompressor(engine string, gsPath, qpdfPath string) (PDFCompressor, error) {
	switch engine {
	case "mock", "dev":
		return NewDevMockCompressor(), nil
	case "qpdf":
		comp := NewQPDFCompressor(qpdfPath)
		return comp, nil
	case "ghostscript":
		comp := NewGhostscriptCompressor(gsPath)
		return comp, nil
	case "auto", "":
		// Try Ghostscript first, then QPDF
		gs := NewGhostscriptCompressor(gsPath)
		if gs.IsAvailable() {
			return gs, nil
		}
		qpdf := NewQPDFCompressor(qpdfPath)
		if qpdf.IsAvailable() {
			return qpdf, nil
		}
		// Return Ghostscript by default
		return gs, nil
	default:
		return nil, fmt.Errorf("unknown compressor engine '%s', supported: ghostscript, qpdf, mock, auto", engine)
	}
}
