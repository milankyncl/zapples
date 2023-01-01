package storage

import "errors"

type CreateFeatureData struct {
	Key         string
	Description *string
}

type UpdateFeatureData struct {
	Key         string
	Description *string
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
