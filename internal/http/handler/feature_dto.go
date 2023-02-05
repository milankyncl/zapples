package handler

import (
	"github.com/milankyncl/feature-toggles/internal/storage"
	"time"
)

type FeatureDto struct {
	Id           int     `json:"id"`
	Key          string  `json:"key"`
	Description  *string `json:"description"`
	Enabled      bool    `json:"enabled"`
	EnabledSince *string `json:"enabled_since"`
	EnabledUntil *string `json:"enabled_until"`
	CreatedAt    string  `json:"created_at"`
}

func featureToDto(f storage.Feature) FeatureDto {
	var enabledSinceStr *string
	if enabledSinceStr = nil; f.EnabledSince != nil {
		formattedEnabledSince := f.EnabledSince.Format(time.RFC3339Nano)
		enabledSinceStr = &formattedEnabledSince
	}
	var enabledUntilStr *string
	if enabledUntilStr = nil; f.EnabledUntil != nil {
		formattedUntilSince := f.EnabledUntil.Format(time.RFC3339Nano)
		enabledUntilStr = &formattedUntilSince
	}
	return FeatureDto{
		Id:           f.Id,
		Key:          f.Key,
		Description:  f.Description,
		Enabled:      f.Enabled,
		EnabledSince: enabledSinceStr,
		EnabledUntil: enabledUntilStr,
		CreatedAt:    f.CreatedAt.Format(time.RFC3339),
	}
}
