package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drinkingandcoding/twotop-backend/database"
	"github.com/drinkingandcoding/twotop-backend/rest/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func getRecipes(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	var recipes []database.Recipe
	err = db.Find(&recipes).Error
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(recipes)
	if err != nil {
		panic(err)
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

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Mount("/recipes", routes.RecipeResource{}.Routes())

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&database.Recipe{})

	fmt.Println("Starting server")
	servErr := http.ListenAndServe(":3000", r)
	if servErr != nil {
		panic(servErr)
	}
}
