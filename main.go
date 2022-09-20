package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/drinkingandcoding/twotop-backend/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	dsn := "host=localhost user=twotop password=twotop dbname=twotop port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&database.Recipe{})
	db.Create(&database.Recipe{
		Name:        "Hot Chicken",
		Keywords:    "hot, chicken",
		Description: "Of the Nashville variety",
		Url:         "https://hot.chicken",
		Yield:       4,
		Ingredients: "Chicken, Heat",
		Steps: []database.Instruction{{
			Name: "Fry chicken",
			Step: "1",
		}, {
			Name: "Cayenne death",
			Step: "2",
		},
		},
	})

	var chicken database.Recipe
	db.First(&chicken, "name = ?", "Hot Chicken")

	fmt.Printf("Hot chicken: %+v", chicken)

	fmt.Println("Starting server")
	servErr := http.ListenAndServe(":3000", r)
	if servErr != nil {
		panic(servErr)
	}
}
