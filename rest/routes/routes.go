package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/drinkingandcoding/twotop-backend/database"
	"github.com/go-chi/chi/v5"
)

const recipePath = "/recipes/"

type RecipeResource struct{}

func (rr RecipeResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", getRecipes)
	r.Post("/", createRecipe)
	r.Route("/{recipeID}", func(r chi.Router) {
		r.Get("/", getRecipeByID)   // GET /recipes/123
		r.Delete("/", deleteRecipe) // DELETE /recipes/123
	})
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

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	recipe := &database.Recipe{}
	var recipeID string
	if strings.HasPrefix(r.URL.Path, recipePath) {
		recipeID = r.URL.Path[len(recipePath):]
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Delete(&recipe, recipeID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getRecipeByID(w http.ResponseWriter, r *http.Request) {
	recipe := &database.Recipe{}
	var recipeID string
	if strings.HasPrefix(r.URL.Path, recipePath) {
		recipeID = r.URL.Path[len(recipePath):]
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.First(&recipe, recipeID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

type Filters struct {
	Keywords   []string `json:"keywords"`
	Cuisines   []string `json:"cuisine"`
	Categories []string `json:"category"`
}

func GetFilters(w http.ResponseWriter, r *http.Request) {
	recipes := []database.Recipe{}
	// db query for keywords, cuisines, category
	validFilters := Filters{}
	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Distinct("keywords", "cuisine", "category").Find(&recipes).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, v := range recipes {
		validFilters.Keywords = append(validFilters.Keywords, v.Keywords...)
		validFilters.Categories = append(validFilters.Categories, v.Category)
		validFilters.Cuisines = append(validFilters.Cuisines, v.Cuisine)
	}
	response, err := json.Marshal(validFilters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
