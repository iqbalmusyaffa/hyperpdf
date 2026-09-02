package dto

import (
	"time"

	"cobagolang/backend/internal/model"

	"github.com/google/uuid"
)

// CompressRequest represents the form data payload for compression
type CompressRequest struct {
	CompressionLevel model.CompressionLevel `form:"compression_level" validate:"required,oneof=LOW MEDIUM HIGH ULTRA_EXTREME EXTREME RECOMMENDED HIGH_FIDELITY STUDIO_MASTER CUSTOM_TARGET"`
}

// JobResponse represents the sanitized public job payload
type JobResponse struct {
	ID                    uuid.UUID              `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	OriginalFilename      string                 `json:"original_filename" example:"document.pdf"`
	OriginalSize          int64                  `json:"original_size" example:"5242880"`
	CompressedSize        int64                  `json:"compressed_size" example:"1835008"`
	SavedBytes            int64                  `json:"saved_bytes" example:"3407872"`
	CompressionPercentage float64                `json:"compression_percentage" example:"65.0"`
	CompressionLevel      model.CompressionLevel `json:"compression_level" example:"MEDIUM"`
	Status                model.JobStatus        `json:"status" example:"COMPLETED"`
	ErrorMessage          string                 `json:"error_message,omitempty" example:""`
	CreatedAt             time.Time              `json:"created_at" example:"2026-09-02T14:30:00Z"`
	CompletedAt           *time.Time             `json:"completed_at,omitempty" example:"2026-09-02T14:30:05Z"`
}

// ToJobResponse converts a model.PDFJob to a JobResponse DTO
func ToJobResponse(job *model.PDFJob) JobResponse {
	return JobResponse{
		ID:                    job.ID,
		OriginalFilename:      job.OriginalFilename,
		OriginalSize:          job.OriginalSize,
		CompressedSize:        job.CompressedSize,
		SavedBytes:            job.SavedBytes,
		CompressionPercentage: job.CompressionPercentage,
		CompressionLevel:      job.CompressionLevel,
		Status:                job.Status,
		ErrorMessage:          job.ErrorMessage,
		CreatedAt:             job.CreatedAt,
		CompletedAt:           job.CompletedAt,
	}
}
