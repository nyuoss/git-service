package main

import (
	"fmt"
	"log"
	"net/http"

	git_service "go-template/git_functions"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// Test GET API endpoint
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		git_service.TestEndpoint(w, r)
	})

	// GET API endpoint to search commit by message
	router.HandleFunc("/getCommitsByMessage", func(w http.ResponseWriter, r *http.Request) {
		git_service.GetCommitsByMessage(w, r)
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	port := 8000
	fmt.Printf("Server is running on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
