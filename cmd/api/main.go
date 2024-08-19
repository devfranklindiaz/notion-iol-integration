package main

import (
	"log"
	"net/http"
    "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/devfranklindiaz/notion-iol-integration/domain/service"
    "github.com/devfranklindiaz/notion-iol-integration/infrastructure/api"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env  %v", err)
	}
	log.Println("Environment loaded")

    notionService := service.NewNotionService()
    log.Println("Notion service created")

    notionHandler := api.NewNotionHandler(notionService)
    log.Println("Notion handler created")

    r := mux.NewRouter()

    apiRouter := r.PathPrefix("/api/v1").Subrouter()
    apiRouter.HandleFunc("/connect", notionHandler.Connect).Methods(http.MethodPost)
    log.Println("API router created")

    err = http.ListenAndServe(":8081", r)

    if err != nil {
        log.Fatalf("Error starting server %v", err)
    }
}
