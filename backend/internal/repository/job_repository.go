package repository

import (
	"context"
	"errors"
	"fmt"

	"cobagolang/backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrJobNotFound = errors.New("pdf job not found")
)

// JobRepository defines database operations for PDF jobs
type JobRepository interface {
	Create(ctx context.Context, job *model.PDFJob) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.PDFJob, error)
	Update(ctx context.Context, job *model.PDFJob) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// GormJobRepository implements JobRepository using GORM
type GormJobRepository struct {
	db *gorm.DB
}

// NewJobRepository creates a new GormJobRepository instance
func NewJobRepository(db *gorm.DB) JobRepository {
	return &GormJobRepository{db: db}
}

// Create inserts a new PDF job into the database
func (r *GormJobRepository) Create(ctx context.Context, job *model.PDFJob) error {
	if err := r.db.WithContext(ctx).Create(job).Error; err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}
	return nil
}

// GetByID retrieves a PDF job by its UUID
func (r *GormJobRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.PDFJob, error) {
	var job model.PDFJob
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&job).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrJobNotFound
		}
		return nil, fmt.Errorf("failed to get job by id: %w", err)
	}
	return &job, nil
}

// Update updates an existing PDF job
func (r *GormJobRepository) Update(ctx context.Context, job *model.PDFJob) error {
	result := r.db.WithContext(ctx).Save(job)
	if result.Error != nil {
		return fmt.Errorf("failed to update job: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrJobNotFound
	}
	return nil
}

// Delete removes a job record from the database
func (r *GormJobRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.PDFJob{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete job: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrJobNotFound
	}
	return nil
}
