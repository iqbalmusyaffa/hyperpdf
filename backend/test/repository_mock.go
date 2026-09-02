package test

import (
	"context"
	"sync"

	"cobagolang/backend/internal/model"
	"cobagolang/backend/internal/repository"

	"github.com/google/uuid"
)

// MockJobRepository implements repository.JobRepository in-memory for unit testing
type MockJobRepository struct {
	mu   sync.RWMutex
	jobs map[uuid.UUID]*model.PDFJob
}

func NewMockJobRepository() *MockJobRepository {
	return &MockJobRepository{
		jobs: make(map[uuid.UUID]*model.PDFJob),
	}
}

func (m *MockJobRepository) Create(ctx context.Context, job *model.PDFJob) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.jobs[job.ID] = job
	return nil
}

func (m *MockJobRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.PDFJob, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	job, exists := m.jobs[id]
	if !exists {
		return nil, repository.ErrJobNotFound
	}
	return job, nil
}

func (m *MockJobRepository) Update(ctx context.Context, job *model.PDFJob) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.jobs[job.ID]; !exists {
		return repository.ErrJobNotFound
	}
	m.jobs[job.ID] = job
	return nil
}

func (m *MockJobRepository) Delete(ctx context.Context, id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.jobs[id]; !exists {
		return repository.ErrJobNotFound
	}
	delete(m.jobs, id)
	return nil
}
