package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cobagolang/backend/internal/compressor"
	"cobagolang/backend/internal/config"
	"cobagolang/backend/internal/dto"
	"cobagolang/backend/internal/model"
	"cobagolang/backend/internal/repository"
	"cobagolang/backend/internal/utils"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	ErrFileNotFound     = errors.New("requested file does not exist on disk")
	ErrJobNotCompleted  = errors.New("compression job is not completed yet")
	ErrJobFailedStatus  = errors.New("compression job failed")
)

// PDFService orchestrates PDF upload, compression, merge, split, and job lifecycle
type PDFService interface {
	CompressPDF(ctx context.Context, fileHeader *multipart.FileHeader, level model.CompressionLevel) (*dto.JobResponse, error)
	MergePDF(ctx context.Context, fileHeaders []*multipart.FileHeader) (*dto.MergeResponse, error)
	SplitPDF(ctx context.Context, fileHeader *multipart.FileHeader, splitMode string, pageRanges string) (*dto.SplitResponse, error)
	GetJobByID(ctx context.Context, id uuid.UUID) (*dto.JobResponse, error)
	GetDownloadFile(ctx context.Context, id uuid.UUID) (filePath string, downloadFilename string, err error)
	DeleteJob(ctx context.Context, id uuid.UUID) error
}

// DefaultPDFService implements PDFService
type DefaultPDFService struct {
	repo       repository.JobRepository
	compressor compressor.PDFCompressor
	cfg        *config.Config
	log        *logrus.Logger
}

// NewPDFService creates a new instance of DefaultPDFService
func NewPDFService(
	repo repository.JobRepository,
	comp compressor.PDFCompressor,
	cfg *config.Config,
) PDFService {
	return &DefaultPDFService{
		repo:       repo,
		compressor: comp,
		cfg:        cfg,
		log:        utils.GetLogger(),
	}
}

// CompressPDF handles validation, disk storage, engine execution, metrics and DB persistence
func (s *DefaultPDFService) CompressPDF(ctx context.Context, fileHeader *multipart.FileHeader, level model.CompressionLevel) (*dto.JobResponse, error) {
	// 1. Validate file
	maxBytes := int64(s.cfg.MaxFileSizeMB) * 1024 * 1024
	if err := utils.ValidatePDFFile(fileHeader, maxBytes); err != nil {
		return nil, err
	}

	jobID := uuid.New()
	sanitizedName := utils.SanitizeFilename(fileHeader.Filename)
	originalSize := fileHeader.Size

	s.log.WithFields(logrus.Fields{
		"job_id":            jobID.String(),
		"original_filename": sanitizedName,
		"original_size":     originalSize,
		"level":             level,
	}).Info("Starting PDF upload and compression workflow")

	// 2. Save uploaded file to storage/uploads/<uuid>.pdf
	inputPath := filepath.Join(s.cfg.StorageUploadDir, fmt.Sprintf("%s_input.pdf", jobID.String()))
	src, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		utils.RemoveFileSafely(inputPath)
		return nil, fmt.Errorf("failed to save uploaded file: %w", err)
	}

	// 3. Create initial Job record in PostgreSQL
	now := time.Now().UTC()
	job := &model.PDFJob{
		ID:               jobID,
		OriginalFilename: sanitizedName,
		OriginalSize:     originalSize,
		CompressionLevel: level,
		InputPath:        inputPath,
		Status:           model.StatusProcessing,
		CreatedAt:        now,
	}

	if err := s.repo.Create(ctx, job); err != nil {
		utils.RemoveFileSafely(inputPath)
		s.log.WithError(err).Error("Database error creating job record")
		return nil, fmt.Errorf("failed to create job in database: %w", err)
	}

	// 4. Execute PDF compression with timeout
	outputPath := filepath.Join(s.cfg.StorageOutputDir, fmt.Sprintf("%s_compressed.pdf", jobID.String()))
	job.OutputPath = outputPath

	timeout := time.Duration(s.cfg.CompressTimeoutSeconds) * time.Second
	compCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	s.log.WithField("job_id", jobID.String()).Info("Executing compression engine")
	compErr := s.compressor.Compress(compCtx, inputPath, outputPath, level)

	// Clean up uploaded temporary input file regardless of success to save disk
	utils.RemoveFileSafely(inputPath)

	if compErr != nil {
		s.log.WithFields(logrus.Fields{
			"job_id": jobID.String(),
			"error":  compErr.Error(),
		}).Error("PDF compression failed")

		job.Status = model.StatusFailed
		job.ErrorMessage = "Compression processing failed"
		completedTime := time.Now().UTC()
		job.CompletedAt = &completedTime
		_ = s.repo.Update(ctx, job)

		return nil, fmt.Errorf("compression failed: %w", compErr)
	}

	// 5. Calculate results
	outInfo, err := os.Stat(outputPath)
	if err != nil {
		job.Status = model.StatusFailed
		job.ErrorMessage = "Failed to inspect compressed output file"
		_ = s.repo.Update(ctx, job)
		return nil, fmt.Errorf("failed to read output file metadata: %w", err)
	}

	compressedSize := outInfo.Size()
	savedBytes, percentage := utils.CalculateSavings(originalSize, compressedSize)

	job.CompressedSize = compressedSize
	job.SavedBytes = savedBytes
	job.CompressionPercentage = percentage
	job.Status = model.StatusCompleted
	completedTime := time.Now().UTC()
	job.CompletedAt = &completedTime

	if err := s.repo.Update(ctx, job); err != nil {
		s.log.WithError(err).Error("Failed to update completed job record in database")
		return nil, fmt.Errorf("failed to save compression results: %w", err)
	}

	s.log.WithFields(logrus.Fields{
		"job_id":                 jobID.String(),
		"original_size":          originalSize,
		"compressed_size":        compressedSize,
		"saved_bytes":            savedBytes,
		"compression_percentage": percentage,
	}).Info("PDF compression completed successfully")

	res := dto.ToJobResponse(job)
	return &res, nil
}

// GetJobByID retrieves job metadata by ID
func (s *DefaultPDFService) GetJobByID(ctx context.Context, id uuid.UUID) (*dto.JobResponse, error) {
	job, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	res := dto.ToJobResponse(job)
	return &res, nil
}

// GetDownloadFile prepares safe file download path and filename
func (s *DefaultPDFService) GetDownloadFile(ctx context.Context, id uuid.UUID) (filePath string, downloadFilename string, err error) {
	job, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return "", "", err
	}

	if job.Status == model.StatusFailed {
		return "", "", ErrJobFailedStatus
	}
	if job.Status != model.StatusCompleted {
		return "", "", ErrJobNotCompleted
	}

	if job.OutputPath == "" {
		return "", "", ErrFileNotFound
	}

	if _, err := os.Stat(job.OutputPath); os.IsNotExist(err) {
		return "", "", ErrFileNotFound
	}

	s.log.WithField("job_id", id.String()).Info("Download requested for completed job")
	downloadName := fmt.Sprintf("compressed_%s", job.OriginalFilename)
	return job.OutputPath, downloadName, nil
}

// DeleteJob removes job record and associated storage files
func (s *DefaultPDFService) DeleteJob(ctx context.Context, id uuid.UUID) error {
	job, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Remove storage files safely
	utils.RemoveFileSafely(job.InputPath)
	utils.RemoveFileSafely(job.OutputPath)

	s.log.WithField("job_id", id.String()).Info("Deleted job and cleaned associated files")
	return s.repo.Delete(ctx, id)
}

// MergePDF merges multiple uploaded PDF files in sequence
func (s *DefaultPDFService) MergePDF(ctx context.Context, fileHeaders []*multipart.FileHeader) (*dto.MergeResponse, error) {
	if len(fileHeaders) < 2 {
		return nil, fmt.Errorf("at least 2 PDF files are required to merge")
	}

	maxBytes := int64(s.cfg.MaxFileSizeMB) * 1024 * 1024
	jobID := uuid.New()
	var inputPaths []string
	var totalInputSize int64

	for i, fh := range fileHeaders {
		if err := utils.ValidatePDFFile(fh, maxBytes); err != nil {
			// Clean up already written files
			for _, p := range inputPaths {
				utils.RemoveFileSafely(p)
			}
			return nil, fmt.Errorf("file %d (%s) is invalid: %w", i+1, fh.Filename, err)
		}

		inPath := filepath.Join(s.cfg.StorageUploadDir, fmt.Sprintf("%s_merge_%d.pdf", jobID.String(), i+1))
		src, err := fh.Open()
		if err != nil {
			for _, p := range inputPaths {
				utils.RemoveFileSafely(p)
			}
			return nil, fmt.Errorf("failed to open file %s: %w", fh.Filename, err)
		}

		dst, err := os.Create(inPath)
		if err != nil {
			src.Close()
			for _, p := range inputPaths {
				utils.RemoveFileSafely(p)
			}
			return nil, fmt.Errorf("failed to create temp file for %s: %w", fh.Filename, err)
		}

		_, copyErr := io.Copy(dst, src)
		src.Close()
		dst.Close()
		if copyErr != nil {
			for _, p := range inputPaths {
				utils.RemoveFileSafely(p)
			}
			return nil, fmt.Errorf("failed to write temp file for %s: %w", fh.Filename, copyErr)
		}

		inputPaths = append(inputPaths, inPath)
		totalInputSize += fh.Size
	}

	outputPath := filepath.Join(s.cfg.StorageOutputDir, fmt.Sprintf("%s_merged.pdf", jobID.String()))

	timeout := time.Duration(s.cfg.CompressTimeoutSeconds) * time.Second
	mergeCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	s.log.WithFields(logrus.Fields{
		"job_id":      jobID.String(),
		"files_count": len(fileHeaders),
		"total_size":  totalInputSize,
	}).Info("Starting PDF merge operation")

	err := s.compressor.Merge(mergeCtx, inputPaths, outputPath)

	// Clean temp input files
	for _, p := range inputPaths {
		utils.RemoveFileSafely(p)
	}

	if err != nil {
		s.log.WithError(err).Error("PDF merge failed")
		return nil, fmt.Errorf("pdf merge failed: %w", err)
	}

	outInfo, err := os.Stat(outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read merged file output: %w", err)
	}

	now := time.Now().UTC()
	job := &model.PDFJob{
		ID:                    jobID,
		OriginalFilename:      "merged_document.pdf",
		OriginalSize:          totalInputSize,
		CompressedSize:        outInfo.Size(),
		CompressionLevel:      model.LevelMedium,
		InputPath:             "",
		OutputPath:            outputPath,
		Status:                model.StatusCompleted,
		CreatedAt:             now,
		CompletedAt:           &now,
	}
	if s.repo != nil {
		_ = s.repo.Create(ctx, job)
	}

	res := &dto.MergeResponse{
		ID:             jobID,
		MergedFilename: "merged_document.pdf",
		FileCount:      len(fileHeaders),
		TotalSize:      outInfo.Size(),
		DownloadURL:    fmt.Sprintf("/api/v1/pdf/jobs/%s/download", jobID.String()),
		CreatedAt:      now,
	}

	return res, nil
}

// SplitPDF splits a single PDF file into page ranges or multiple files
func (s *DefaultPDFService) SplitPDF(ctx context.Context, fileHeader *multipart.FileHeader, splitMode string, pageRanges string) (*dto.SplitResponse, error) {
	maxBytes := int64(s.cfg.MaxFileSizeMB) * 1024 * 1024
	if err := utils.ValidatePDFFile(fileHeader, maxBytes); err != nil {
		return nil, err
	}

	jobID := uuid.New()
	sanitizedName := utils.SanitizeFilename(fileHeader.Filename)
	inputPath := filepath.Join(s.cfg.StorageUploadDir, fmt.Sprintf("%s_split_input.pdf", jobID.String()))

	src, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		utils.RemoveFileSafely(inputPath)
		return nil, fmt.Errorf("failed to save temp file: %w", err)
	}

	splitDir := filepath.Join(s.cfg.StorageOutputDir, fmt.Sprintf("%s_split_temp", jobID.String()))

	timeout := time.Duration(s.cfg.CompressTimeoutSeconds) * time.Second
	splitCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	s.log.WithFields(logrus.Fields{
		"job_id":     jobID.String(),
		"split_mode": splitMode,
		"ranges":     pageRanges,
	}).Info("Starting PDF split operation")

	files, err := s.compressor.Split(splitCtx, inputPath, splitDir, splitMode, pageRanges)
	utils.RemoveFileSafely(inputPath)

	if err != nil {
		_ = os.RemoveAll(splitDir)
		s.log.WithError(err).Error("PDF split failed")
		return nil, fmt.Errorf("pdf split failed: %w", err)
	}

	var finalOutputPath string
	var downloadFilename string
	var isZip bool

	if len(files) == 1 {
		// Single extracted PDF
		finalOutputPath = filepath.Join(s.cfg.StorageOutputDir, fmt.Sprintf("%s_extracted.pdf", jobID.String()))
		_ = os.Rename(files[0], finalOutputPath)
		_ = os.RemoveAll(splitDir)
		downloadFilename = fmt.Sprintf("extracted_%s", sanitizedName)
		isZip = false
	} else {
		// Multiple pages -> Zip archive
		finalOutputPath = filepath.Join(s.cfg.StorageOutputDir, fmt.Sprintf("%s_split_pages.zip", jobID.String()))
		if err := utils.CreateZipArchive(files, finalOutputPath); err != nil {
			_ = os.RemoveAll(splitDir)
			return nil, fmt.Errorf("failed to create zip archive of split pages: %w", err)
		}
		_ = os.RemoveAll(splitDir)
		downloadFilename = fmt.Sprintf("split_pages_%s.zip", strings.TrimSuffix(sanitizedName, ".pdf"))
		isZip = true
	}

	now := time.Now().UTC()
	outInfo, _ := os.Stat(finalOutputPath)
	outSize := int64(0)
	if outInfo != nil {
		outSize = outInfo.Size()
	}

	job := &model.PDFJob{
		ID:               jobID,
		OriginalFilename: downloadFilename,
		OriginalSize:     fileHeader.Size,
		CompressedSize:   outSize,
		CompressionLevel: model.LevelMedium,
		OutputPath:       finalOutputPath,
		Status:           model.StatusCompleted,
		CreatedAt:        now,
		CompletedAt:      &now,
	}
	if s.repo != nil {
		_ = s.repo.Create(ctx, job)
	}

	res := &dto.SplitResponse{
		ID:               jobID,
		OriginalFilename: sanitizedName,
		SplitMode:        splitMode,
		PageRanges:       pageRanges,
		GeneratedCount:   len(files),
		IsZipArchive:     isZip,
		DownloadFilename: downloadFilename,
		DownloadURL:      fmt.Sprintf("/api/v1/pdf/jobs/%s/download", jobID.String()),
		CreatedAt:        now,
	}

	return res, nil
}
