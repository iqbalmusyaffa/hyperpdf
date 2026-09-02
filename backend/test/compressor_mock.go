package test

import (
	"context"
	"fmt"
	"os"

	"cobagolang/backend/internal/model"
)

// MockCompressor implements compressor.PDFCompressor for testing
type MockCompressor struct {
	Available      bool
	ShouldFail     bool
	EngineName     string
	CompressedData []byte
}

func NewMockCompressor(available bool, shouldFail bool) *MockCompressor {
	return &MockCompressor{
		Available:      available,
		ShouldFail:     shouldFail,
		EngineName:     "mock_compressor",
		CompressedData: []byte("%PDF-1.4\n%mock compressed content\n%%EOF"),
	}
}

func (m *MockCompressor) Compress(ctx context.Context, inputPath, outputPath string, level model.CompressionLevel) error {
	if m.ShouldFail {
		return fmt.Errorf("mock compression failure")
	}

	// Write mock compressed PDF content
	return os.WriteFile(outputPath, m.CompressedData, 0644)
}

func (m *MockCompressor) Name() string {
	return m.EngineName
}

func (m *MockCompressor) IsAvailable() bool {
	return m.Available
}

func (m *MockCompressor) GetBinaryPath() string {
	return "/usr/bin/mock_engine"
}

func (m *MockCompressor) Merge(ctx context.Context, inputPaths []string, outputPath string) error {
	if m.ShouldFail {
		return fmt.Errorf("mock merge failure")
	}
	return os.WriteFile(outputPath, []byte("%PDF-1.4\n%mock merged\n%%EOF"), 0644)
}

func (m *MockCompressor) Split(ctx context.Context, inputPath, outputDir string, splitMode string, pageRanges string) ([]string, error) {
	if m.ShouldFail {
		return nil, fmt.Errorf("mock split failure")
	}
	outPath := outputDir + "/page_1.pdf"
	_ = os.MkdirAll(outputDir, 0755)
	_ = os.WriteFile(outPath, []byte("%PDF-1.4\n%mock split page\n%%EOF"), 0644)
	return []string{outPath}, nil
}
