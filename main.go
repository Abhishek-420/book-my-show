package main

import (
	"book-my-show/handler"
	"book-my-show/repository"
	"book-my-show/routes"
	"book-my-show/service"
	"log"
	"net/http"
)

func main() {

	repo := repository.NewBookMyShowRepo()
	service := service.NewBookMyShowService(repo)
	handler := handler.NewBookMyShowHandler(service)
	mux := routes.RegisterRoutes(handler)

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
