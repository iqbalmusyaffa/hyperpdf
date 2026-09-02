package compressor

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"cobagolang/backend/internal/model"
)

// DevMockCompressor is a development fallback compressor that simulates compression
type DevMockCompressor struct{}

// NewDevMockCompressor creates a new DevMockCompressor instance
func NewDevMockCompressor() *DevMockCompressor {
	return &DevMockCompressor{}
}

func (m *DevMockCompressor) Name() string {
	return "mock_dev"
}

func (m *DevMockCompressor) IsAvailable() bool {
	return true
}

func (m *DevMockCompressor) GetBinaryPath() string {
	return "builtin://mock-simulator"
}

func (m *DevMockCompressor) Compress(ctx context.Context, inputPath, outputPath string, level model.CompressionLevel) error {
	// Read original file stats
	info, err := os.Stat(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file metadata: %w", err)
	}

	src, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}
	defer src.Close()

	// Determine reduction ratio based on level
	var targetRatio float64
	switch level {
	case model.LevelLow:
		targetRatio = 0.70 // 30% reduction
	case model.LevelMedium:
		targetRatio = 0.40 // 60% reduction
	case model.LevelHigh:
		targetRatio = 0.20 // 80% reduction
	default:
		targetRatio = 0.50
	}

	targetBytes := int64(float64(info.Size()) * targetRatio)
	if targetBytes < 512 {
		targetBytes = info.Size()
	}

	// Read input bytes up to targetBytes, ensuring PDF header and trailer
	buf := make([]byte, targetBytes)
	_, _ = io.ReadFull(src, buf)

	// Ensure valid PDF trailer if truncated
	if len(buf) > 10 {
		copy(buf[:5], []byte("%PDF-"))
		copy(buf[len(buf)-6:], []byte("%%EOF\n"))
	}

	dst, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer dst.Close()

	if _, err := dst.Write(buf); err != nil {
		return fmt.Errorf("failed to write compressed file: %w", err)
	}

	return nil
}

func (m *DevMockCompressor) Merge(ctx context.Context, inputPaths []string, outputPath string) error {
	dst, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create merged file: %w", err)
	}
	defer dst.Close()

	_, _ = dst.WriteString("%PDF-1.4\n% Merged PDF Document\n")
	for i, inPath := range inputPaths {
		_, _ = dst.WriteString(fmt.Sprintf("%% Start File %d: %s\n", i+1, inPath))
		data, err := os.ReadFile(inPath)
		if err == nil {
			_, _ = dst.Write(data)
		}
	}
	_, _ = dst.WriteString("\n%%EOF\n")
	return nil
}

func (m *DevMockCompressor) Split(ctx context.Context, inputPath, outputDir string, splitMode string, pageRanges string) ([]string, error) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create split directory: %w", err)
	}

	var generated []string
	if splitMode == "range" && pageRanges != "" {
		outPath := filepath.Join(outputDir, fmt.Sprintf("extracted_pages_%s.pdf", pageRanges))
		content := []byte(fmt.Sprintf("%%PDF-1.4\n%% Extracted Pages [%s] from %s\n%%%%EOF\n", pageRanges, inputPath))
		if err := os.WriteFile(outPath, content, 0644); err != nil {
			return nil, err
		}
		generated = append(generated, outPath)
	} else {
		// Create 3 split pages by default
		for i := 1; i <= 3; i++ {
			outPath := filepath.Join(outputDir, fmt.Sprintf("page_%03d.pdf", i))
			content := []byte(fmt.Sprintf("%%PDF-1.4\n%% Page %d from %s\n%%%%EOF\n", i, inputPath))
			if err := os.WriteFile(outPath, content, 0644); err != nil {
				return nil, err
			}
			generated = append(generated, outPath)
		}
	}

	return generated, nil
}
