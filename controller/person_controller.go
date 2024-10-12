package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zelalem-t8/go-crud-challenge/database"
	"github.com/zelalem-t8/go-crud-challenge/model"
)

type PersonController struct {
	DB *database.InMemoryDB
}

// Create handles the creation of a new person.
func (pc *PersonController) Create(w http.ResponseWriter, r *http.Request) {
	var person model.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input fields
	if person.Name == "" || person.Age <= 0 || len(person.Hobbies) == 0 {
		http.Error(w, "Invalid person data", http.StatusBadRequest)
		return
	}

	createdPerson := pc.DB.Create(person)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPerson)
}

// GetAll handles retrieving all persons.
func (pc *PersonController) GetAll(w http.ResponseWriter, r *http.Request) {
	persons := pc.DB.GetAll()
	json.NewEncoder(w).Encode(persons)
}

// Get handles retrieving a person by ID.
func (pc *PersonController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["personId"]

	person, exists := pc.DB.Get(id)
	if !exists {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(person)
}

// Update handles updating an existing person.
func (pc *PersonController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["personId"]

	var updatedPerson model.Person
	if err := json.NewDecoder(r.Body).Decode(&updatedPerson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input fields
	if updatedPerson.Name == "" || updatedPerson.Age <= 0 || len(updatedPerson.Hobbies) == 0 {
		http.Error(w, "Invalid person data", http.StatusBadRequest)
		return
	}

	person, success := pc.DB.Update(id, updatedPerson)
	if !success {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(person)
}

// Delete handles removing a person by ID.
func (pc *PersonController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["personId"]

	if success := pc.DB.Delete(id); !success {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
