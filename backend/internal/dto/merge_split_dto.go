package dto

import (
	"time"

	"github.com/google/uuid"
)

// MergeResponse represents the result of a PDF merge operation
type MergeResponse struct {
	ID             uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	MergedFilename string    `json:"merged_filename" example:"merged_document.pdf"`
	FileCount      int       `json:"file_count" example:"3"`
	TotalSize      int64     `json:"total_size" example:"3145728"`
	DownloadURL    string    `json:"download_url" example:"/api/v1/pdf/jobs/550e8400-e29b-41d4-a716-446655440000/download"`
	CreatedAt      time.Time `json:"created_at"`
}

// SplitResponse represents the result of a PDF split operation
type SplitResponse struct {
	ID               uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	OriginalFilename string    `json:"original_filename" example:"contract.pdf"`
	SplitMode        string    `json:"split_mode" example:"range"`
	PageRanges       string    `json:"page_ranges,omitempty" example:"1-3"`
	GeneratedCount   int       `json:"generated_count" example:"1"`
	IsZipArchive     bool      `json:"is_zip_archive" example:"false"`
	DownloadFilename string    `json:"download_filename" example:"split_contract.pdf"`
	DownloadURL      string    `json:"download_url" example:"/api/v1/pdf/jobs/550e8400-e29b-41d4-a716-446655440000/download"`
	CreatedAt        time.Time `json:"created_at"`
}
