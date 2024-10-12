package database

import (
	"sync"

	"github.com/google/uuid"
	"github.com/zelalem-t8/go-crud-challenge/model"
)

// InMemoryDB simulates an in-memory database for Person objects.
type InMemoryDB struct {
	data map[string]model.Person
	mu   sync.Mutex
}

// NewInMemoryDB initializes a new in-memory database.
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]model.Person),
	}
}

// Create adds a new person to the database.
func (db *InMemoryDB) Create(person model.Person) model.Person {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Generate a UUID using the google/uuid library
	person.ID = uuid.New().String()
	db.data[person.ID] = person
	return person
}

// Get retrieves a person by ID.
func (db *InMemoryDB) Get(id string) (model.Person, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	person, exists := db.data[id]
	return person, exists
}

// GetAll retrieves all persons from the database.
func (db *InMemoryDB) GetAll() []model.Person {
	db.mu.Lock()
	defer db.mu.Unlock()
	persons := make([]model.Person, 0, len(db.data))
	for _, person := range db.data {
		persons = append(persons, person)
	}
	return persons
}

// Update modifies an existing person in the database.
func (db *InMemoryDB) Update(id string, updatedPerson model.Person) (model.Person, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	if person, exists := db.data[id]; exists {
		updatedPerson.ID = person.ID
		db.data[id] = updatedPerson
		return updatedPerson, true
	}
	return model.Person{}, false
}

// Delete removes a person from the database.
func (db *InMemoryDB) Delete(id string) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, exists := db.data[id]; exists {
		delete(db.data, id)
		return true
	}
	return false
}
