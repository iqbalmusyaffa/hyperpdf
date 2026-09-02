package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// JobStatus represents the state of a PDF compression job
type JobStatus string

const (
	StatusPending    JobStatus = "PENDING"
	StatusProcessing JobStatus = "PROCESSING"
	StatusCompleted  JobStatus = "COMPLETED"
	StatusFailed     JobStatus = "FAILED"
)

// CompressionLevel represents the level of compression requested
type CompressionLevel string

const (
	LevelLow    CompressionLevel = "LOW"
	LevelMedium CompressionLevel = "MEDIUM"
	LevelHigh   CompressionLevel = "HIGH"
)

// PDFJob represents a PDF compression job database record
type PDFJob struct {
	ID                    uuid.UUID        `gorm:"type:uuid;primaryKey" json:"id"`
	OriginalFilename      string           `gorm:"type:varchar(255);not null" json:"original_filename"`
	OriginalSize          int64            `gorm:"not null" json:"original_size"`
	CompressedSize        int64            `gorm:"default:0" json:"compressed_size"`
	SavedBytes            int64            `gorm:"default:0" json:"saved_bytes"`
	CompressionPercentage float64          `gorm:"type:numeric(5,2);default:0.00" json:"compression_percentage"`
	CompressionLevel      CompressionLevel `gorm:"type:varchar(20);not null" json:"compression_level"`
	InputPath             string           `gorm:"type:text;not null" json:"input_path"`
	OutputPath            string           `gorm:"type:text" json:"output_path"`
	Status                JobStatus        `gorm:"type:varchar(20);not null;default:'PENDING';index" json:"status"`
	ErrorMessage          string           `gorm:"type:text" json:"error_message,omitempty"`
	CreatedAt             time.Time        `gorm:"not null;default:CURRENT_TIMESTAMP;index" json:"created_at"`
	CompletedAt           *time.Time       `json:"completed_at,omitempty"`
}

// TableName specifies custom table name
func (PDFJob) TableName() string {
	return "pdf_jobs"
}

// BeforeCreate GORM hook to ensure UUID is generated
func (job *PDFJob) BeforeCreate(tx *gorm.DB) error {
	if job.ID == uuid.Nil {
		job.ID = uuid.New()
	}
	return nil
}
