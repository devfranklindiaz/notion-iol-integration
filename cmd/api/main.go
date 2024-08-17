package main

import (
	"log"
	"net/http"

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
	notionHandler := api.NewNotionHandler(notionService)

	http.HandleFunc("/connect", notionHandler.Connect)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
