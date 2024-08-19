package database

import "time"

// BaseModel defines common field behaviour.
type BaseModel interface {
	GetID() int
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

// Model defines common fields for all models.
type Model struct {
	Id        int       `json:"id" gorm:"primary_key;autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// GetID provides the id for the model.
func (m Model) GetID() int {
	return m.Id
}

// GetCreatedAt provides the creation date for the model.
func (m Model) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// GetUpdatedAt provides the update time of the model.
func (m Model) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}
