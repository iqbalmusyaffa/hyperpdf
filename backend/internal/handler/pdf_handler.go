package handler

import (
	"errors"
	"strings"

	"cobagolang/backend/internal/dto"
	"cobagolang/backend/internal/model"
	"cobagolang/backend/internal/repository"
	"cobagolang/backend/internal/service"
	"cobagolang/backend/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PDFHandler handles PDF compression, status retrieval, download, and deletion
type PDFHandler struct {
	pdfService service.PDFService
}

// NewPDFHandler creates a new PDFHandler instance
func NewPDFHandler(pdfService service.PDFService) *PDFHandler {
	return &PDFHandler{pdfService: pdfService}
}

// CompressPDF handles PDF upload and executes compression
// @Summary Compress PDF
// @Description Upload a PDF file and compress it based on the specified level (LOW, MEDIUM, HIGH)
// @Tags PDF
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "PDF File to compress"
// @Param compression_level formData string false "Compression Level (LOW, MEDIUM, HIGH)" default(MEDIUM)
// @Success 200 {object} dto.APIResponse{data=dto.JobResponse}
// @Failure 400 {object} dto.APIErrorResponse
// @Failure 500 {object} dto.APIErrorResponse
// @Router /api/v1/pdf/compress [post]
func (h *PDFHandler) CompressPDF(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.NewErrorResponse("No PDF file uploaded", "Form field 'file' is required"),
		)
	}

	rawLevel := c.FormValue("compression_level", "MEDIUM")
	level := model.CompressionLevel(strings.ToUpper(strings.TrimSpace(rawLevel)))

	req := dto.CompressRequest{
		CompressionLevel: level,
	}

	if valErrors := utils.ValidateStruct(&req); len(valErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.NewErrorResponse("Validation failed", valErrors...),
		)
	}

	jobResponse, err := h.pdfService.CompressPDF(c.Context(), fileHeader, level)
	if err != nil {
		if errors.Is(err, utils.ErrInvalidExtension) ||
			errors.Is(err, utils.ErrInvalidMimeType) ||
			errors.Is(err, utils.ErrInvalidPDFHeader) ||
			errors.Is(err, utils.ErrEmptyFile) ||
			errors.Is(err, utils.ErrFileTooLarge) {
			return c.Status(fiber.StatusBadRequest).JSON(
				dto.NewErrorResponse("Invalid PDF file", err.Error()),
			)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.NewErrorResponse("Failed to compress PDF", err.Error()),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		dto.NewSuccessResponse("PDF compressed successfully", jobResponse),
	)
}

// GetJob retrieves job status and metadata by ID
// @Summary Get Job Details
// @Description Retrieve compression job details and status by Job UUID
// @Tags PDF
// @Produce json
// @Param id path string true "Job UUID"
// @Success 200 {object} dto.APIResponse{data=dto.JobResponse}
// @Failure 400 {object} dto.APIErrorResponse
// @Failure 404 {object} dto.APIErrorResponse
// @Router /api/v1/pdf/jobs/{id} [get]
func (h *PDFHandler) GetJob(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.NewErrorResponse("Invalid Job ID", "ID must be a valid UUID"),
		)
	}

	jobResponse, err := h.pdfService.GetJobByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrJobNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(
				dto.NewErrorResponse("Job not found", "No job exists with the given ID"),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.NewErrorResponse("Database error", "Failed to retrieve job details"),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		dto.NewSuccessResponse("Job retrieved successfully", jobResponse),
	)
}

// DownloadCompressedPDF serves the compressed PDF file
// @Summary Download Compressed PDF
// @Description Download the compressed PDF file for a completed job
// @Tags PDF
// @Produce application/pdf
// @Param id path string true "Job UUID"
// @Success 200 {file} binary
// @Failure 400 {object} dto.APIErrorResponse
// @Failure 404 {object} dto.APIErrorResponse
// @Failure 500 {object} dto.APIErrorResponse
// @Router /api/v1/pdf/jobs/{id}/download [get]
func (h *PDFHandler) DownloadCompressedPDF(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.NewErrorResponse("Invalid Job ID", "ID must be a valid UUID"),
		)
	}

	filePath, downloadName, err := h.pdfService.GetDownloadFile(c.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrJobNotFound) || errors.Is(err, service.ErrFileNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(
				dto.NewErrorResponse("File not found", "The requested compressed file is not available"),
			)
		}
		if errors.Is(err, service.ErrJobNotCompleted) {
			return c.Status(fiber.StatusBadRequest).JSON(
				dto.NewErrorResponse("Job not ready", "Compression job is still in progress"),
			)
		}
		if errors.Is(err, service.ErrJobFailedStatus) {
			return c.Status(fiber.StatusBadRequest).JSON(
				dto.NewErrorResponse("Job failed", "Cannot download file for a failed compression job"),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.NewErrorResponse("Download failed", "An error occurred while preparing file download"),
		)
	}

	c.Set("Content-Type", "application/pdf")
	return c.Download(filePath, downloadName)
}

// DeleteJob deletes a job and its associated storage files
// @Summary Delete Job
// @Description Delete a compression job and all corresponding files from disk
// @Tags PDF
// @Produce json
// @Param id path string true "Job UUID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.APIErrorResponse
// @Failure 404 {object} dto.APIErrorResponse
// @Router /api/v1/pdf/jobs/{id} [delete]
func (h *PDFHandler) DeleteJob(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.NewErrorResponse("Invalid Job ID", "ID must be a valid UUID"),
		)
	}

	if err := h.pdfService.DeleteJob(c.Context(), id); err != nil {
		if errors.Is(err, repository.ErrJobNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(
				dto.NewErrorResponse("Job not found", "No job exists with the given ID"),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.NewErrorResponse("Deletion failed", "Failed to delete job"),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		dto.NewSuccessResponse("Job and associated files deleted successfully", nil),
	)
}

// MergePDF merges multiple uploaded PDF files
// @Summary Merge multiple PDF files
// @Description Upload multiple PDF files and combine them in order
// @Tags PDF
// @Accept multipart/form-data
// @Produce json
// @Param files formData file true "PDF files to merge"
// @Success 200 {object} dto.APIResponse{data=dto.MergeResponse}
// @Failure 400 {object} dto.APIErrorResponse
// @Failure 500 {object} dto.APIErrorResponse
// @Router /api/v1/pdf/merge [post]
func (h *PDFHandler) MergePDF(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.NewErrorResponse("Invalid form data", "Failed to parse multipart form"),
		)
	}

	files := form.File["files"]
	if len(files) < 2 {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.NewErrorResponse("Insufficient files", "Please upload at least 2 PDF files to merge"),
		)
	}

	mergeRes, err := h.pdfService.MergePDF(c.Context(), files)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.NewErrorResponse("Failed to merge PDF files", err.Error()),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		dto.NewSuccessResponse("PDF files merged successfully", mergeRes),
	)
}

// SplitPDF splits a single PDF file into pages or range
// @Summary Split PDF file
// @Description Split PDF by page range or extract every page to a ZIP archive
// @Tags PDF
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "PDF file to split"
// @Param split_mode formData string false "Split mode (range or all)" default(all)
// @Param page_ranges formData string false "Page ranges (e.g. 1-3 or 2)"
// @Success 200 {object} dto.APIResponse{data=dto.SplitResponse}
// @Failure 400 {object} dto.APIErrorResponse
// @Failure 500 {object} dto.APIErrorResponse
// @Router /api/v1/pdf/split [post]
func (h *PDFHandler) SplitPDF(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			dto.NewErrorResponse("No PDF file uploaded", "Form field 'file' is required"),
		)
	}

	splitMode := strings.ToLower(strings.TrimSpace(c.FormValue("split_mode", "all")))
	pageRanges := strings.TrimSpace(c.FormValue("page_ranges", ""))

	splitRes, err := h.pdfService.SplitPDF(c.Context(), fileHeader, splitMode, pageRanges)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dto.NewErrorResponse("Failed to split PDF file", err.Error()),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		dto.NewSuccessResponse("PDF file split successfully", splitRes),
	)
}
