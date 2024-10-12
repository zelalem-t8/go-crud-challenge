package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zelalem-t8/go-crud-challenge/controller"
	"github.com/zelalem-t8/go-crud-challenge/database"
	"github.com/zelalem-t8/go-crud-challenge/model"
)

// TestCreatePerson tests the Create method
func TestCreate(t *testing.T) {
	db := database.NewInMemoryDB()
	personController := &controller.PersonController{DB: db}

	payload := map[string]interface{}{
		"name":    "John Doe",
		"age":     30,
		"hobbies": []string{"reading", "swimming"},
	}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/person", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(personController.Create)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdPerson model.Person
	err = json.NewDecoder(rr.Body).Decode(&createdPerson)
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", createdPerson.Name)
	assert.Equal(t, 30, createdPerson.Age)
	assert.ElementsMatch(t, []string{"reading", "swimming"}, createdPerson.Hobbies)
}

// TestGetAllPersons tests the GetAll method
func TestGetAll(t *testing.T) {
	db := database.NewInMemoryDB()
	personController := &controller.PersonController{DB: db}

	db.Create(model.Person{Name: "Jane Doe", Age: 25, Hobbies: []string{"painting"}})
	db.Create(model.Person{Name: "John Doe", Age: 30, Hobbies: []string{"reading"}})

	req, err := http.NewRequest("GET", "/person", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(personController.GetAll)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var persons []model.Person
	err = json.NewDecoder(rr.Body).Decode(&persons)
	assert.Nil(t, err)
	assert.Len(t, persons, 2)
}

// TestGetPerson tests the Get method
func TestGet(t *testing.T) {
	db := database.NewInMemoryDB()
	personController := &controller.PersonController{DB: db}

	person := db.Create(model.Person{Name: "Jane Doe", Age: 25, Hobbies: []string{"painting"}})

	req, err := http.NewRequest("GET", "/person/"+person.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/person/{personId}", personController.Get)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var returnedPerson model.Person
	err = json.NewDecoder(rr.Body).Decode(&returnedPerson)
	assert.Nil(t, err)
	assert.Equal(t, "Jane Doe", returnedPerson.Name)
	assert.Equal(t, 25, returnedPerson.Age)
	assert.ElementsMatch(t, []string{"painting"}, returnedPerson.Hobbies)
}

// TestUpdatePerson tests the Update method
func TestUpdate(t *testing.T) {
	db := database.NewInMemoryDB()
	personController := &controller.PersonController{DB: db}

	person := db.Create(model.Person{Name: "Mark Doe", Age: 40, Hobbies: []string{"hiking"}})

	updatedPayload := map[string]interface{}{
		"name":    "Mark Updated",
		"age":     45,
		"hobbies": []string{"traveling"},
	}
	payloadBytes, _ := json.Marshal(updatedPayload)
	req, err := http.NewRequest("PUT", "/person/"+person.ID, bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/person/{personId}", personController.Update)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var updatedPerson model.Person
	err = json.NewDecoder(rr.Body).Decode(&updatedPerson)
	assert.Nil(t, err)
	assert.Equal(t, "Mark Updated", updatedPerson.Name)
	assert.Equal(t, 45, updatedPerson.Age)
	assert.ElementsMatch(t, []string{"traveling"}, updatedPerson.Hobbies)
}

// TestDeletePerson tests the Delete method
func TestDelete(t *testing.T) {
	db := database.NewInMemoryDB()
	personController := &controller.PersonController{DB: db}

	person := db.Create(model.Person{Name: "Mark Doe", Age: 40, Hobbies: []string{"hiking"}})

	req, err := http.NewRequest("DELETE", "/person/"+person.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/person/{personId}", personController.Delete)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	_, exists := db.Get(person.ID)
	assert.False(t, exists)
}

// TestInvalidPersonID tests invalid personId access
func TestInvalidPersonID(t *testing.T) {
	db := database.NewInMemoryDB()
	personController := &controller.PersonController{DB: db}

	req, err := http.NewRequest("GET", "/person/invalid-id", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/person/{personId}", personController.Get)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), "Person not found")
}
