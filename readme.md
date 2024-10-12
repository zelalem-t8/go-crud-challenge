# Go CRUD API for Person Management

This project implements a simple CRUD (Create, Read, Update, Delete) API for managing Person records using Go and an in-memory database. The API provides endpoints to create, retrieve, update, and delete Person objects, as well as import and export functionality for CSV files.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Usage](#usage)
- [Testing](#testing)
- [Importing and Exporting CSV](#importing-and-exporting-csv)
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
   cd <repository-directory>
   ```

2. **Install dependencies:**

Make sure you have Go installed on your system. Run the following command to download the necessary packages:

```go get -u github.com/gorilla/mux
go get -u github.com/google/uuid
```

3. **Run the application:**

To start the server, run:`bash go run main.go`

4. **Access the API:**
   _Create Person_

POST /person
Request Body:

```{
  "name": "John Doe",
  "age": 30,
  "hobbies": ["reading", "gaming"]
}
```

_Get All Persons_

GET /person

_Get Person by ID_

GET /person/{personId}

_Update Person_

PUT /person/{personId}

Request Body

```{
  "name": "John Doe Updated",
  "age": 31,
  "hobbies": ["reading", "traveling"]
}
```
