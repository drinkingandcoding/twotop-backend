package routes

import (
	"encoding/json"
	"net/http"

	"github.com/drinkingandcoding/twotop-backend/database"
	"github.com/go-chi/chi/v5"
)

type RecipeResource struct{}

func (rr RecipeResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", getRecipes)
	r.Post("/", createRecipe)
	return r
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var recipes []database.Recipe
	err = db.Find(&recipes).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(recipes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func createRecipe(w http.ResponseWriter, r *http.Request) {
	recipe := &database.Recipe{}
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = db.Create(recipe).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
