package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/milankyncl/feature-toggles/internal/storage"
	"log"
	"net/http"
	"strconv"
	"time"
)

type GetFeaturesResponseDto struct {
	Data []FeatureDto `json:"data"`
}

func GetFeaturesHandler(storage storage.Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		features, err := storage.GetAll()
		if err != nil {
			log.Println(err)
			http.Error(w, "Unexpected error happened", http.StatusInternalServerError)
			return
		}

		dtos := make([]FeatureDto, 0)
		for _, f := range features {
			dtos = append(dtos, featureToDto(f))
		}

		b, err := json.Marshal(&GetFeaturesResponseDto{
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

type GetFeatureResponseDto struct {
	Data FeatureDto `json:"data"`
}

func GetFeatureHandler(adapter storage.Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		feat, err := adapter.GetOne(id)
		if err == storage.ErrFeatureNotFound {
			http.Error(w, "Feature not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, "Unexpected error happened", http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(&GetFeatureResponseDto{
			Data: featureToDto(feat),
		})
		if err != nil {
			http.Error(w, "Unexpected error happened", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
	}
}

type ManageFeatureRequestDto struct {
	Key          string  `json:"key"`
	Description  *string `json:"description"`
	EnabledSince *string `json:"enabledSince"`
	EnabledUntil *string `json:"enabledUntil"`
}

func CreateFeatureHandler(adapter storage.Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto ManageFeatureRequestDto
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		since, until, err := parseFeatureEnableDates(dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = adapter.Create(storage.CreateFeatureData{
			Key:          dto.Key,
			Description:  dto.Description,
			EnabledSince: since,
			EnabledUntil: until,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateFeatureHandler(adapter storage.Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var dto ManageFeatureRequestDto
		err = json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		since, until, err := parseFeatureEnableDates(dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = adapter.Update(id, storage.UpdateFeatureData{
			Key:          dto.Key,
			Description:  dto.Description,
			EnabledSince: since,
			EnabledUntil: until,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

type ToggleFeatureRequestDto struct {
	Enabled bool `json:"enabled"`
}

func ToggleFeatureHandler(adapter storage.Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var dto ToggleFeatureRequestDto
		err = json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = adapter.Toggle(id, dto.Enabled)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteFeatureHandler(adapter storage.Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = adapter.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func parseFeatureEnableDates(dto ManageFeatureRequestDto) (*time.Time, *time.Time, error) {
	var enabledSince *time.Time
	if dto.EnabledSince != nil {
		parsedEnabledSince, err := time.Parse(time.RFC3339Nano, *dto.EnabledSince)
		if err != nil {
			return nil, nil, err
		}
		enabledSince = &parsedEnabledSince
	}
	var enabledUntil *time.Time
	if dto.EnabledUntil != nil {
		parsedEnabledUntil, err := time.Parse(time.RFC3339Nano, *dto.EnabledUntil)
		if err != nil {
			return nil, nil, err
		}
		enabledUntil = &parsedEnabledUntil
	}
	return enabledSince, enabledUntil, nil
}
