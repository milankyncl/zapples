package handler

import (
	"encoding/json"
	"github.com/milankyncl/feature-toggles/internal/storage"
	"log"
	"net/http"
)

type GetFeatureTogglesResponseDto struct {
	Data []FeatureToggleDto `json:"data"`
}

func GetFeatureTogglesHandler(storage storage.Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		features, err := storage.GetAll()
		if err != nil {
			log.Println(err)
			http.Error(w, "Unexpected error happened", http.StatusInternalServerError)
			return
		}

		dtos := make([]FeatureToggleDto, 0)
		for _, f := range features {
			dtos = append(dtos, featureToToggleDto(f))
		}

		b, err := json.Marshal(&GetFeatureTogglesResponseDto{
			Data: dtos,
		})
		if err != nil {
			log.Println(err)
			http.Error(w, "Unexpected error happened", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
	}
}
