package main

import (
	"fmt"
	"net/http"

	"github.com/drinkingandcoding/twotop-backend/database"
	"github.com/drinkingandcoding/twotop-backend/rest/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

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
	db.AutoMigrate(&database.Recipe{}, &database.Author{}, &database.Instruction{}, &database.Recipe{})

	fmt.Println("Running server")
	servErr := http.ListenAndServe(":3000", r)
	if servErr != nil {
		panic(servErr)
	}
}
