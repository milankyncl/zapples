package storage

import (
	"errors"
	"time"
)

type CreateFeatureData struct {
	Key          string
	Description  *string
	EnabledSince *time.Time
	EnabledUntil *time.Time
}

type UpdateFeatureData struct {
	Key          string
	Description  *string
	EnabledSince *time.Time
	EnabledUntil *time.Time
}

var (
	ErrFeatureNotFound = errors.New("feature not found")
)

type Adapter interface {
	GetAll() ([]Feature, error)
	GetOne(id int) (Feature, error)
	Create(data CreateFeatureData) error
	Update(id int, data UpdateFeatureData) error
	Toggle(id int, enabled bool) error
	Delete(id int) error
}
