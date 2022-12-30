package handler

import (
	"github.com/milankyncl/feature-toggles/internal/storage"
	"time"
)

type FeatureDto struct {
	Id          int     `json:"id"`
	Key         string  `json:"key"`
	Description *string `json:"description"`
	Enabled     bool    `json:"enabled"`
	CreatedAt   string  `json:"created_at"`
}

func featureToDto(f storage.Feature) FeatureDto {
	return FeatureDto{
		Id:          f.Id,
		Key:         f.Key,
		Description: f.Description,
		Enabled:     f.Enabled,
		CreatedAt:   f.CreatedAt.Format(time.RFC3339),
	}
}
