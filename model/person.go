package model

// Person represents a person object with ID, Name, Age, and Hobbies.
type Person struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}
