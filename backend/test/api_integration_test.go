package test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"net/textproto"
	"os"
	"path/filepath"
	"testing"

	"cobagolang/backend/internal/config"
	"cobagolang/backend/internal/dto"
	"cobagolang/backend/internal/handler"
	"cobagolang/backend/internal/middleware"
	"cobagolang/backend/internal/model"
	"cobagolang/backend/internal/repository"
	"cobagolang/backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestApp(t *testing.T) (*fiber.App, *MockJobRepository, *MockCompressor, *config.Config) {
	tempDir := filepath.Join(os.TempDir(), "pdf_test_"+uuid.New().String())
	uploadDir := filepath.Join(tempDir, "uploads")
	outputDir := filepath.Join(tempDir, "outputs")
	require.NoError(t, os.MkdirAll(uploadDir, 0755))
	require.NoError(t, os.MkdirAll(outputDir, 0755))

	t.Cleanup(func() {
		// Give Windows filesystem a brief moment to release handles if fasthttp is closing
		_ = os.RemoveAll(tempDir)
	})

	cfg := &config.Config{
		StorageUploadDir:       uploadDir,
		StorageOutputDir:       outputDir,
		MaxFileSizeMB:          50,
		CompressTimeoutSeconds: 5,
		RateLimitMax:           1000,
	}

	repo := NewMockJobRepository()
	comp := NewMockCompressor(true, false)
	svc := service.NewPDFService(repo, comp, cfg)
	pdfHandler := handler.NewPDFHandler(svc)

	app := fiber.New()
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	api := app.Group("/api/v1")
	pdf := api.Group("/pdf")
	{
		pdf.Post("/compress", pdfHandler.CompressPDF)
		pdf.Post("/merge", pdfHandler.MergePDF)
		pdf.Post("/split", pdfHandler.SplitPDF)
		pdf.Get("/jobs/:id", pdfHandler.GetJob)
		pdf.Get("/jobs/:id/download", pdfHandler.DownloadCompressedPDF)
		pdf.Delete("/jobs/:id", pdfHandler.DeleteJob)
	}

	return app, repo, comp, cfg
}

func TestAPI_CompressPDF_Endpoint(t *testing.T) {
	app, _, _, _ := setupTestApp(t)

	// Create multipart body
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	partHeader := make(textproto.MIMEHeader)
	partHeader.Set("Content-Disposition", `form-data; name="file"; filename="sample.pdf"`)
	partHeader.Set("Content-Type", "application/pdf")
	part, err := writer.CreatePart(partHeader)
	require.NoError(t, err)

	_, err = part.Write([]byte("%PDF-1.4\n%integration test content\n%%EOF"))
	require.NoError(t, err)

	_ = writer.WriteField("compression_level", "MEDIUM")
	_ = writer.Close()

	req := httptest.NewRequest("POST", "/api/v1/pdf/compress", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req, 10000)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	respBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var apiRes struct {
		Success bool            `json:"success"`
		Message string          `json:"message"`
		Data    dto.JobResponse `json:"data"`
	}
	require.NoError(t, json.Unmarshal(respBytes, &apiRes))
	assert.True(t, apiRes.Success)
	assert.Equal(t, "sample.pdf", apiRes.Data.OriginalFilename)
	assert.Equal(t, model.StatusCompleted, apiRes.Data.Status)
	assert.NotEqual(t, uuid.Nil, apiRes.Data.ID)
}

func TestAPI_GetJob_Endpoint(t *testing.T) {
	app, repo, _, _ := setupTestApp(t)

	testID := uuid.New()
	job := &model.PDFJob{
		ID:                    testID,
		OriginalFilename:      "job_test.pdf",
		OriginalSize:          1024,
		CompressedSize:        512,
		SavedBytes:            512,
		CompressionPercentage: 50.0,
		CompressionLevel:      model.LevelMedium,
		Status:                model.StatusCompleted,
	}
	require.NoError(t, repo.Create(nil, job))

	req := httptest.NewRequest("GET", "/api/v1/pdf/jobs/"+testID.String(), nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAPI_Download_Endpoint(t *testing.T) {
	app, repo, _, cfg := setupTestApp(t)

	testID := uuid.New()
	outputPath := filepath.Join(cfg.StorageOutputDir, testID.String()+"_compressed.pdf")
	require.NoError(t, os.WriteFile(outputPath, []byte("%PDF-1.4\n%downloadable content\n%%EOF"), 0644))

	job := &model.PDFJob{
		ID:               testID,
		OriginalFilename: "document.pdf",
		OutputPath:       outputPath,
		Status:           model.StatusCompleted,
	}
	require.NoError(t, repo.Create(nil, job))

	req := httptest.NewRequest("GET", "/api/v1/pdf/jobs/"+testID.String()+"/download", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	downloaded, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.True(t, bytes.HasPrefix(downloaded, []byte("%PDF-")))
}

func TestAPI_Delete_Endpoint(t *testing.T) {
	app, repo, _, cfg := setupTestApp(t)

	testID := uuid.New()
	outputPath := filepath.Join(cfg.StorageOutputDir, testID.String()+"_compressed.pdf")
	require.NoError(t, os.WriteFile(outputPath, []byte("%PDF-1.4\n%%EOF"), 0644))

	job := &model.PDFJob{
		ID:         testID,
		OutputPath: outputPath,
		Status:     model.StatusCompleted,
	}
	require.NoError(t, repo.Create(nil, job))

	req := httptest.NewRequest("DELETE", "/api/v1/pdf/jobs/"+testID.String(), nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)

	// Verify deletion
	_, err = repo.GetByID(nil, testID)
	assert.ErrorIs(t, err, repository.ErrJobNotFound)
}

func TestAPI_MergePDF_Endpoint(t *testing.T) {
	app, _, _, _ := setupTestApp(t)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Add file 1
	h1 := make(textproto.MIMEHeader)
	h1.Set("Content-Disposition", `form-data; name="files"; filename="doc1.pdf"`)
	h1.Set("Content-Type", "application/pdf")
	p1, _ := writer.CreatePart(h1)
	_, _ = p1.Write([]byte("%PDF-1.4\n%part1\n%%EOF"))

	// Add file 2
	h2 := make(textproto.MIMEHeader)
	h2.Set("Content-Disposition", `form-data; name="files"; filename="doc2.pdf"`)
	h2.Set("Content-Type", "application/pdf")
	p2, _ := writer.CreatePart(h2)
	_, _ = p2.Write([]byte("%PDF-1.4\n%part2\n%%EOF"))

	_ = writer.Close()

	req := httptest.NewRequest("POST", "/api/v1/pdf/merge", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req, 10000)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAPI_SplitPDF_Endpoint(t *testing.T) {
	app, _, _, _ := setupTestApp(t)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="file"; filename="sample.pdf"`)
	h.Set("Content-Type", "application/pdf")
	part, _ := writer.CreatePart(h)
	_, _ = part.Write([]byte("%PDF-1.4\n%split content\n%%EOF"))

	_ = writer.WriteField("split_mode", "range")
	_ = writer.WriteField("page_ranges", "1-2")
	_ = writer.Close()

	req := httptest.NewRequest("POST", "/api/v1/pdf/split", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req, 10000)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
}
