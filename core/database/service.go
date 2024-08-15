package database

import (
	"context"
	"gorm.io/gorm"
)

type Service[T any] interface {
	Create(ctx context.Context, entity *T) error
	GetByID(ctx context.Context, id int) (*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, filter func(*gorm.DB) *gorm.DB) ([]T, error)
}

// service is the implementation of the generic Service.
type service[T any] struct {
	repo Repository[T]
}

// NewService creates a new generic service.
func NewService[T any](repo Repository[T]) Service[T] {
	return &service[T]{repo: repo}
}

// Create creates a new record.
func (s *service[T]) Create(ctx context.Context, entity *T) error {
	return s.repo.Create(ctx, entity)
}

// GetByID retrieves a record by its ID.
func (s *service[T]) GetByID(ctx context.Context, id int) (*T, error) {
	return s.repo.FindByID(ctx, id)
}

// Update updates an existing record.
func (s *service[T]) Update(ctx context.Context, entity *T) error {
	return s.repo.Update(ctx, entity)
}

// Delete deletes a record by its ID.
func (s *service[T]) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// List lists all records with an optional filter.
func (s *service[T]) List(ctx context.Context, filter func(*gorm.DB) *gorm.DB) ([]T, error) {
	return s.repo.List(ctx, filter)
}
