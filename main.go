package main

import (
	"fmt"
	"log"
	"net/http"

	gitService "git-service/git_functions"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// serve swagger-ui
	sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./pkg/swagger-ui")))
	router.PathPrefix("/swaggerui/").Handler(sh)

	// Test GET API endpoint
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		gitService.TestEndpoint(w, r)
	})

	router.HandleFunc(
		"/v1/{owner}/{repo}/branch/getActiveBranches",
		gitService.GetActiveBranches).
		Methods(http.MethodGet)

	router.HandleFunc(
		"/v1/{owner}/{repo}/branch/getBranchByTag",
		gitService.GetBranchByTag).
		Methods(http.MethodGet)

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
