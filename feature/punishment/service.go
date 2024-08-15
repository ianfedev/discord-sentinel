package punishment

import (
	"context"
	"discord-sentinel/core/database"
	"gorm.io/gorm"
)

// Service defines the interface for Punishment operations.
type Service interface {
	Create(ctx context.Context, punishment *Punishment) error
	GetByID(ctx context.Context, id int) (*Punishment, error)
	Update(ctx context.Context, punishment *Punishment) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, filter func(*gorm.DB) *gorm.DB) ([]Punishment, error)
}

// service is the implementation of the Punishment Service.
type service struct {
	repo database.Repository[Punishment]
}

// NewService creates a new Punishment service.
func NewService(repo database.Repository[Punishment]) Service {
	return &service{repo: repo}
}

// Create creates a new punishment record.
func (s *service) Create(ctx context.Context, punishment *Punishment) error {
	return s.repo.Create(ctx, punishment)
}

// GetByID retrieves a punishment record by its ID.
func (s *service) GetByID(ctx context.Context, id int) (*Punishment, error) {
	return s.repo.FindByID(ctx, id)
}

// Update updates an existing punishment record.
func (s *service) Update(ctx context.Context, punishment *Punishment) error {
	return s.repo.Update(ctx, punishment)
}

// Delete deletes a punishment record by its ID.
func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// List lists all punishments with an optional filter.
func (s *service) List(ctx context.Context, filter func(*gorm.DB) *gorm.DB) ([]Punishment, error) {
	return s.repo.List(ctx, filter)
}
