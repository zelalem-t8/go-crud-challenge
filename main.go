package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zelalem-t8/go-crud-challenge/controller"
	"github.com/zelalem-t8/go-crud-challenge/database"
)

func main() {
	// Initialize the in-memory database.
	db := database.NewInMemoryDB()
	pc := &controller.PersonController{DB: db}

	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/person", pc.Create).Methods("POST")
	r.HandleFunc("/person", pc.GetAll).Methods("GET")
	r.HandleFunc("/person/{personId}", pc.Get).Methods("GET")
	r.HandleFunc("/person/{personId}", pc.Update).Methods("PUT")
	r.HandleFunc("/person/{personId}", pc.Delete).Methods("DELETE")

	// Handle non-existing routes
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Endpoint not found", http.StatusNotFound)
	})

	// Add CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all domains for simplicity
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
