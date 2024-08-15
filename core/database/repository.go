package database

import (
	"context"
	"gorm.io/gorm"
)

// Repository defines an interface with common CRUD behaviour for
// database connection.
type Repository[T any] interface {
	Create(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id int) (*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, query func(*gorm.DB) *gorm.DB) ([]T, error)
}

// GormRepository implements Repository interface using GORM.
type GormRepository[T any] struct {
	db *gorm.DB
}

// NewGormRepository creates a new GormRepository.
func wNewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

// Create inserts a new record.
func (r *GormRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

// FindByID retrieves a record by its ID.
func (r *GormRepository[T]) FindByID(ctx context.Context, id int) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Update updates an existing record.
func (r *GormRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

// Delete removes a record by its ID.
func (r *GormRepository[T]) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(new(T), id).Error
}

// List retrieves all records with an optional query customization.
func (r *GormRepository[T]) List(ctx context.Context, query func(*gorm.DB) *gorm.DB) ([]T, error) {
	var entities []T
	db := r.db.WithContext(ctx)
	if query != nil {
		db = query(db)
	}
	if err := db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}
