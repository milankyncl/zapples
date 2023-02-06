package handler

import (
	"github.com/milankyncl/feature-toggles/internal/storage"
	"time"
)

type FeatureToggleDto struct {
	Key     string `json:"key"`
	Toggled bool   `json:"toggled"`
}

func featureToToggleDto(feature storage.Feature) FeatureToggleDto {
	tgld := feature.Enabled
	if tgld && feature.EnabledSince != nil && time.Now().Before(*feature.EnabledSince) {
		tgld = false
	}
	if tgld && feature.EnabledUntil != nil && time.Now().After(*feature.EnabledUntil) {
		tgld = false
	}
	return FeatureToggleDto{
		Key:     feature.Key,
		Toggled: tgld,
	}
}
