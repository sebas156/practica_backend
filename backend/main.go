package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sebas156/backend/handlers"
	"github.com/sebas156/backend/repository"
)

func main() {
	db := repository.NewDataBase()

	handler := handlers.NewHandlerCharacter(db)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/superhero", handler.GetHeroByName())

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
