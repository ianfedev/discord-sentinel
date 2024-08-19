package database

import (
	"context"
	"gorm.io/gorm"
	"reflect"
)

// Repository implements Repository interface using GORM.
type Repository[T any] struct {
	db *gorm.DB
}

// NewGormRepository creates a new Repository.
func NewGormRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db: db}
}

// Create inserts a new record.
func (r *Repository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

// FindByID retrieves a record by its ID.
func (r *Repository[T]) FindByID(ctx context.Context, id int) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Update updates an existing record, ensuring ID and CreatedAt are not modified and UpdatedAt is refreshed.
func (r *Repository[T]) Update(ctx context.Context, entity *T) error {

	baseEntity, ok := any(entity).(BaseModel)
	if !ok {
		return gorm.ErrInvalidData
	}

	dbEntity, err := r.FindByID(ctx, baseEntity.GetID())
	if err != nil {
		return err
	}

	entityValue := reflect.ValueOf(entity).Elem()
	idField := entityValue.FieldByName("Id")
	if idField.IsValid() {
		idField.Set(reflect.Zero(idField.Type()))
	}

	updates := r.db.Model(dbEntity).Updates(entity)
	if updates.Error != nil {
		return updates.Error
	}

	return r.db.Model(dbEntity).Update("updated_at", gorm.Expr("NOW()")).Error
}

// Delete removes a record by its ID.
func (r *Repository[T]) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(new(T), id).Error
}

// List retrieves all records with an optional query customization.
func (r *Repository[T]) List(ctx context.Context, query func(*gorm.DB) *gorm.DB) ([]T, error) {
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
