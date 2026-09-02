package test

import (
	"bytes"
	"mime/multipart"
	"net/textproto"
	"testing"

	"cobagolang/backend/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeFilename(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "clean simple filename",
			input:    "report.pdf",
			expected: "report.pdf",
		},
		{
			name:     "path traversal unix",
			input:    "../../../../etc/passwd.pdf",
			expected: "passwd.pdf",
		},
		{
			name:     "path traversal windows",
			input:    `..\..\..\Windows\System32\cmd.pdf`,
			expected: "cmd.pdf",
		},
		{
			name:     "special characters and spaces",
			input:    "My Cool Report & Final (1) [v2].pdf",
			expected: "My_Cool_Report___Final__1___v2_.pdf",
		},
		{
			name:     "missing extension",
			input:    "document",
			expected: "document.pdf",
		},
		{
			name:     "empty filename",
			input:    "",
			expected: "document.pdf",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.SanitizeFilename(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculateSavings(t *testing.T) {
	tests := []struct {
		name           string
		original       int64
		compressed     int64
		wantSaved      int64
		wantPercentage float64
	}{
		{
			name:           "50% compression",
			original:       1000,
			compressed:     500,
			wantSaved:      500,
			wantPercentage: 50.0,
		},
		{
			name:           "65% compression",
			original:       10000,
			compressed:     3500,
			wantSaved:      6500,
			wantPercentage: 65.0,
		},
		{
			name:           "larger compressed file (no savings)",
			original:       1000,
			compressed:     1200,
			wantSaved:      0,
			wantPercentage: 0.0,
		},
		{
			name:           "zero original size",
			original:       0,
			compressed:     0,
			wantSaved:      0,
			wantPercentage: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saved, pct := utils.CalculateSavings(tt.original, tt.compressed)
			assert.Equal(t, tt.wantSaved, saved)
			assert.InDelta(t, tt.wantPercentage, pct, 0.01)
		})
	}
}

func TestFormatBytes(t *testing.T) {
	assert.Equal(t, "500 B", utils.FormatBytes(500))
	assert.Equal(t, "1.00 KB", utils.FormatBytes(1024))
	assert.Equal(t, "5.00 MB", utils.FormatBytes(5*1024*1024))
	assert.Equal(t, "1.50 GB", utils.FormatBytes(1610612736))
}

func TestValidatePDFFile(t *testing.T) {
	t.Run("valid PDF header", func(t *testing.T) {
		content := []byte("%PDF-1.4\n%âãÏÓ\n1 0 obj\n<<>>\nendobj\ntrailer\n<<>>\n%%EOF")
		fh := createMockFileHeader("valid.pdf", content)

		err := utils.ValidatePDFFile(fh, 10*1024*1024)
		assert.NoError(t, err)
	})

	t.Run("invalid extension", func(t *testing.T) {
		content := []byte("%PDF-1.4\ncontent")
		fh := createMockFileHeader("image.png", content)

		err := utils.ValidatePDFFile(fh, 10*1024*1024)
		assert.ErrorIs(t, err, utils.ErrInvalidExtension)
	})

	t.Run("invalid header / not a PDF", func(t *testing.T) {
		content := []byte("PK\x03\x04This is a zip file, not a PDF")
		fh := createMockFileHeader("fake.pdf", content)

		err := utils.ValidatePDFFile(fh, 10*1024*1024)
		assert.ErrorIs(t, err, utils.ErrInvalidPDFHeader)
	})

	t.Run("file exceeds maximum size", func(t *testing.T) {
		content := []byte("%PDF-1.4\n" + string(make([]byte, 2048)))
		fh := createMockFileHeader("large.pdf", content)

		err := utils.ValidatePDFFile(fh, 1024) // 1KB limit
		assert.ErrorIs(t, err, utils.ErrFileTooLarge)
	})
}

func createMockFileHeader(filename string, content []byte) *multipart.FileHeader {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="file"; filename="`+filename+`"`)
	h.Set("Content-Type", "application/pdf")

	part, _ := writer.CreatePart(h)
	_, _ = part.Write(content)
	_ = writer.Close()

	reader := multipart.NewReader(&buf, writer.Boundary())
	form, _ := reader.ReadForm(int64(len(content) + 1024))
	return form.File["file"][0]
}
