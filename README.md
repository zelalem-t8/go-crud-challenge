# Go CRUD API for Person Management

This project implements a simple CRUD (Create, Read, Update, Delete) API for managing Person records using Go and an in-memory database. The API provides endpoints to create, retrieve, update, and delete Person objects, as well as import and export functionality for CSV files.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
- [Access the API](#Access-the-API)
- [Importing CSV](#importing-csv)
- [Error Handling](#error-handling)
- [License](#license)

## Features

- CRUD operations for Person records
- In-memory database implementation
- CSV import and export functionality
- Basic error handling and validation
- CORS support for frontend applications

## Technologies

- Go (Golang)
- Gorilla Mux (for routing)
- UUID (for unique ID generation)

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/zelalem-t8/go-crud-challenge.git
   cd go-crud-challenge
   ```

2. **Install dependencies:**

Make sure you have Go installed on your system. Run the following command to download the necessary packages:

```go get -u github.com/gorilla/mux
go get -u github.com/google/uuid
```

3. **Run the application:**

To start the server, run:`bash go run main.go`

## Access the API

_Create Person_

POST /person
Request Body:

```{
  "name": "John Doe",
  "age": 30,
  "hobbies": ["reading", "gaming"]
}
```

Get All Persons `GET /person`

Get Person by ID `GET /person/{personId}`

Update Person `PUT /person/{personId}`

Request Body

```{
  "name": "John Doe Updated",
  "age": 31,
  "hobbies": ["reading", "traveling"]
}
```
Delete Person ;`DELETE /person/{personId}`

## CSV Import/Export Endpoints
Import Persons from CSV

POST /person/import
Form Data:
File: CSV file containing persons data with headers: name, age, hobbies.


