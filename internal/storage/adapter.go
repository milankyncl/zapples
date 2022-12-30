package storage

type CreateFeatureData struct {
	Key         string
	Description *string
}

type UpdateFeatureData struct {
	Key         string
	Description *string
}

type Adapter interface {
	GetAll() ([]Feature, error)
	Create(data CreateFeatureData) error
	Update(id int, data UpdateFeatureData) error
	Toggle(id int, enabled bool) error
}
