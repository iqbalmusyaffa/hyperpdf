package test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"cobagolang/backend/internal/config"
	"cobagolang/backend/internal/model"
	"cobagolang/backend/internal/repository"
	"cobagolang/backend/internal/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestService(t *testing.T, shouldFailComp bool) (service.PDFService, *MockJobRepository, *config.Config) {
	tempDir := t.TempDir()
	uploadDir := filepath.Join(tempDir, "uploads")
	outputDir := filepath.Join(tempDir, "outputs")
	require.NoError(t, os.MkdirAll(uploadDir, 0755))
	require.NoError(t, os.MkdirAll(outputDir, 0755))

	cfg := &config.Config{
		StorageUploadDir:       uploadDir,
		StorageOutputDir:       outputDir,
		MaxFileSizeMB:          50,
		CompressTimeoutSeconds: 5,
	}

	repo := NewMockJobRepository()
	comp := NewMockCompressor(true, shouldFailComp)
	svc := service.NewPDFService(repo, comp, cfg)

	return svc, repo, cfg
}

func TestPDFService_CompressPDF_Success(t *testing.T) {
	svc, repo, _ := setupTestService(t, false)
	ctx := context.Background()

	rawPDF := []byte("%PDF-1.4\n%test pdf content large enough to compress\n" + string(make([]byte, 1024)) + "\n%%EOF")
	fh := createMockFileHeader("document.pdf", rawPDF)

	res, err := svc.CompressPDF(ctx, fh, model.LevelMedium)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, "document.pdf", res.OriginalFilename)
	assert.Equal(t, int64(len(rawPDF)), res.OriginalSize)
	assert.True(t, res.CompressedSize > 0)
	assert.Equal(t, model.StatusCompleted, res.Status)
	assert.Equal(t, model.LevelMedium, res.CompressionLevel)

	// Verify repo state
	savedJob, err := repo.GetByID(ctx, res.ID)
	require.NoError(t, err)
	assert.Equal(t, model.StatusCompleted, savedJob.Status)
}

func TestPDFService_CompressPDF_EngineFailure(t *testing.T) {
	svc, repo, _ := setupTestService(t, true) // Engine will fail
	ctx := context.Background()

	rawPDF := []byte("%PDF-1.4\n%test pdf\n%%EOF")
	fh := createMockFileHeader("broken.pdf", rawPDF)

	res, err := svc.CompressPDF(ctx, fh, model.LevelHigh)
	assert.Error(t, err)
	assert.Nil(t, res)

	// Verify a failed job was persisted
	assert.Equal(t, 1, len(repo.jobs))
	for _, job := range repo.jobs {
		assert.Equal(t, model.StatusFailed, job.Status)
	}
}

func TestPDFService_GetJobByID(t *testing.T) {
	svc, repo, _ := setupTestService(t, false)
	ctx := context.Background()

	testID := uuid.New()
	job := &model.PDFJob{
		ID:               testID,
		OriginalFilename: "sample.pdf",
		OriginalSize:     2048,
		Status:           model.StatusCompleted,
	}
	require.NoError(t, repo.Create(ctx, job))

	res, err := svc.GetJobByID(ctx, testID)
	require.NoError(t, err)
	assert.Equal(t, testID, res.ID)
	assert.Equal(t, "sample.pdf", res.OriginalFilename)

	// Non-existent ID
	_, err = svc.GetJobByID(ctx, uuid.New())
	assert.ErrorIs(t, err, repository.ErrJobNotFound)
}

func TestPDFService_GetDownloadFile(t *testing.T) {
	svc, repo, cfg := setupTestService(t, false)
	ctx := context.Background()

	testID := uuid.New()
	outputPath := filepath.Join(cfg.StorageOutputDir, testID.String()+"_compressed.pdf")
	require.NoError(t, os.WriteFile(outputPath, []byte("%PDF-1.4\n%%EOF"), 0644))

	job := &model.PDFJob{
		ID:               testID,
		OriginalFilename: "my_doc.pdf",
		OutputPath:       outputPath,
		Status:           model.StatusCompleted,
	}
	require.NoError(t, repo.Create(ctx, job))

	filePath, dlName, err := svc.GetDownloadFile(ctx, testID)
	require.NoError(t, err)
	assert.Equal(t, outputPath, filePath)
	assert.Equal(t, "compressed_my_doc.pdf", dlName)

	// Pending job cannot be downloaded
	pendingID := uuid.New()
	pendingJob := &model.PDFJob{
		ID:     pendingID,
		Status: model.StatusPending,
	}
	require.NoError(t, repo.Create(ctx, pendingJob))
	_, _, err = svc.GetDownloadFile(ctx, pendingID)
	assert.ErrorIs(t, err, service.ErrJobNotCompleted)
}

func TestPDFService_DeleteJob(t *testing.T) {
	svc, repo, cfg := setupTestService(t, false)
	ctx := context.Background()

	testID := uuid.New()
	outputPath := filepath.Join(cfg.StorageOutputDir, testID.String()+"_compressed.pdf")
	require.NoError(t, os.WriteFile(outputPath, []byte("%PDF-1.4\n%%EOF"), 0644))

	job := &model.PDFJob{
		ID:               testID,
		OriginalFilename: "to_delete.pdf",
		OutputPath:       outputPath,
		Status:           model.StatusCompleted,
	}
	require.NoError(t, repo.Create(ctx, job))

	err := svc.DeleteJob(ctx, testID)
	require.NoError(t, err)

	// File should be removed from disk
	_, statErr := os.Stat(outputPath)
	assert.True(t, os.IsNotExist(statErr))

	// Job should not exist in repo
	_, getErr := repo.GetByID(ctx, testID)
	assert.ErrorIs(t, getErr, repository.ErrJobNotFound)
}

func TestPDFService_MergePDF(t *testing.T) {
	svc, _, _ := setupTestService(t, false)
	ctx := context.Background()

	pdf1 := []byte("%PDF-1.4\n%file 1\n%%EOF")
	pdf2 := []byte("%PDF-1.4\n%file 2\n%%EOF")
	fh1 := createMockFileHeader("doc1.pdf", pdf1)
	fh2 := createMockFileHeader("doc2.pdf", pdf2)

	res, err := svc.MergePDF(ctx, []*multipart.FileHeader{fh1, fh2})
	require.NoError(t, err)
	require.NotNil(t, res)
	assert.Equal(t, 2, res.FileCount)
	assert.Equal(t, "merged_document.pdf", res.MergedFilename)
	assert.NotEmpty(t, res.DownloadURL)
}

func TestPDFService_SplitPDF(t *testing.T) {
	svc, _, _ := setupTestService(t, false)
	ctx := context.Background()

	rawPDF := []byte("%PDF-1.4\n%multi-page doc\n%%EOF")
	fh := createMockFileHeader("contract.pdf", rawPDF)

	res, err := svc.SplitPDF(ctx, fh, "range", "1-2")
	require.NoError(t, err)
	require.NotNil(t, res)
	assert.Equal(t, "contract.pdf", res.OriginalFilename)
	assert.Equal(t, "range", res.SplitMode)
	assert.Equal(t, "1-2", res.PageRanges)
}
