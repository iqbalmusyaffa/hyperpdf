package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	ErrEmptyFile        = errors.New("file is empty")
	ErrFileTooLarge     = errors.New("file exceeds maximum allowed size")
	ErrInvalidExtension = errors.New("only .pdf files are allowed")
	ErrInvalidMimeType  = errors.New("file MIME type is not application/pdf")
	ErrInvalidPDFHeader = errors.New("file is not a valid PDF document (missing PDF signature)")
)

// PDFMagicBytes is standard PDF header signature
var PDFMagicBytes = []byte("%PDF-")

var safeFilenameRegex = regexp.MustCompile(`[^a-zA-Z0-9._-]`)

// ValidatePDFFile validates extension, size, MIME type, and magic bytes header
func ValidatePDFFile(fh *multipart.FileHeader, maxSizeBytes int64) error {
	if fh == nil || fh.Size == 0 {
		return ErrEmptyFile
	}

	if fh.Size > maxSizeBytes {
		return fmt.Errorf("%w: max size is %d MB", ErrFileTooLarge, maxSizeBytes/(1024*1024))
	}

	// Validate extension
	ext := strings.ToLower(filepath.Ext(fh.Filename))
	if ext != ".pdf" {
		return ErrInvalidExtension
	}

	file, err := fh.Open()
	if err != nil {
		return fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer file.Close()

	// Read first 512 bytes for MIME detection and magic bytes
	header := make([]byte, 512)
	n, err := file.Read(header)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read file header: %w", err)
	}

	// Check magic bytes "%PDF-"
	if !bytes.HasPrefix(header[:n], PDFMagicBytes) {
		// Some PDF files have leading whitespace or BOM, search first 1024 bytes
		if !bytes.Contains(header[:n], PDFMagicBytes) {
			return ErrInvalidPDFHeader
		}
	}

	// Check detected MIME type
	mimeType := http.DetectContentType(header[:n])
	if mimeType != "application/pdf" && mimeType != "application/octet-stream" {
		// DetectContentType sometimes returns application/octet-stream for PDFs with BOM,
		// but since magic bytes matched, we allow it. Otherwise reject non-PDF MIME.
		return ErrInvalidMimeType
	}

	return nil
}

// SanitizeFilename cleans the filename to prevent path traversal and unsafe characters
func SanitizeFilename(filename string) string {
	trimmed := strings.TrimSpace(filename)
	if trimmed == "" || trimmed == "." || trimmed == ".." {
		return "document.pdf"
	}
	base := filepath.Base(trimmed)
	if base == "." || base == ".." || base == "/" || base == `\` {
		return "document.pdf"
	}
	cleaned := safeFilenameRegex.ReplaceAllString(base, "_")
	if !strings.HasSuffix(strings.ToLower(cleaned), ".pdf") {
		cleaned += ".pdf"
	}
	if cleaned == ".pdf" || cleaned == "_.pdf" || cleaned == "..pdf" {
		cleaned = "document.pdf"
	}
	return cleaned
}

// CalculateSavings computes saved bytes and reduction percentage
func CalculateSavings(originalSize, compressedSize int64) (savedBytes int64, percentage float64) {
	if originalSize <= 0 {
		return 0, 0.0
	}
	if compressedSize < originalSize {
		savedBytes = originalSize - compressedSize
		percentage = (float64(savedBytes) / float64(originalSize)) * 100.0
		// Round to 2 decimal places
		percentage = float64(int(percentage*100+0.5)) / 100.0
	} else {
		savedBytes = 0
		percentage = 0.0
	}
	return savedBytes, percentage
}

// RemoveFileSafely removes a file if it exists, without raising errors if it doesn't
func RemoveFileSafely(path string) {
	if path != "" {
		_ = os.Remove(path)
	}
}

// FormatBytes formats byte sizes into human readable strings (e.g., 5.2 MB)
func FormatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
