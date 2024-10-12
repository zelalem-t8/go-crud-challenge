package controller

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
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

// ImportFromCSV handles the import of persons from a CSV file uploaded by the user.
func (pc *PersonController) ImportFromCSV(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data (set a maximum file size limit, e.g., 10MB)
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	// Get the file from form input
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to retrieve file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	// Skip the header row
	_, err = reader.Read()
	if err != nil {
		http.Error(w, "Error reading CSV header", http.StatusBadRequest)
		return
	}

	var importedPersons []model.Person
	// Iterate through the CSV file and parse each row
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break // End of file
		}
		if err != nil {
			http.Error(w, "Error reading CSV file", http.StatusBadRequest)
			return
		}

		// Map CSV fields to the Person model
		age, err := strconv.Atoi(line[2]) // Convert age to integer
		if err != nil {
			http.Error(w, "Invalid age format", http.StatusBadRequest)
			return
		}

		person := model.Person{
			ID:      uuid.NewString(),             // Generate new ID
			Name:    line[1],                      // Get name from CSV
			Age:     age,                          // Get age from CSV
			Hobbies: strings.Split(line[3], ", "), // Split hobbies by ", "
		}

		// Add the person to the in-memory DB
		pc.DB.Create(person)
		importedPersons = append(importedPersons, person)
	}

	// Return the imported persons
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(importedPersons)
}
