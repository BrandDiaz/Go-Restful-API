package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main(){
	port := "3000"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Printf("Server Starting... http://localhost:%s", port)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, router *http.Request){
		w.Header().Set("Content-Type", "application/json")
	})

	router.Mount("/products", buyersResource{}.Routes())

  	log.Fatal(http.ListenAndServe(":" + port, router))

}