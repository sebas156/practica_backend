package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sebas156/backend/repository"
)

type HandlerCharacter struct {
	DB *repository.DataBase
}

func NewHandlerCharacter(db *repository.DataBase) *HandlerCharacter {
	return &HandlerCharacter{
		DB: db,
	}
}

func (hc *HandlerCharacter) GetHeroByName() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
			return
		}

		heroName := r.URL.Query().Get("hero")

		if heroName == "" {
			http.Error(w, "No name was submitted", http.StatusBadRequest)
			return
		}

		hero, ok := hc.DB.Memory[heroName]

		if !ok {
			http.Error(w, "No hero was found with that name", http.StatusNotFound)
			return
		}

		payload, err := json.Marshal(hero)

		if err != nil {
			http.Error(w, "JSON codification failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)

	})
}
