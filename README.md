# E-Commerce System with Go

A full-featured e-commerce system built in **Go**, using **PostgreSQL** as the database and **Docker** for easy setup.  
Includes user registration & login, product management, and a shopping cart system.

---

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Setup with Docker](#setup-with-docker)
- [Database Configuration](#database-configuration)
- [Run the Application](#run-the-application)
- [API Endpoints](#api-endpoints)
- [Postman Example](#postman-example)

---

## Features

- User registration and login
- Product CRUD (Create, Read, Update, Delete)
- Shopping cart with add, view, clear, and checkout
- PostgreSQL as database
- RESTful API with Gorilla Mux

---

## Tech Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Router:** Gorilla Mux
- **Docker:** Containerization for database

---

## Prerequisites

- Go >= 1.21
- Docker Desktop
- Postman (for testing API)

---

## Setup with Docker

Run PostgreSQL container:

```bash
docker run --name ecommerce-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=ecommerce -p 5432:5432 -d postgres
```

## Go Project Setup

Initialize Go module:

```bash
go mod tidy
```

Install dependencies:

```bash
go get github.com/gorilla/mux
go get github.com/lib/pq
go get github.com/google/uuid
```

## Configuration

In main.go, configure database connection:

```go
connStr := "host=localhost port=5432 user=postgres password=123456 dbname=ecommerce sslmode=disable"
db, err := sql.Open("postgres", connStr)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

## Running the API

Start the server:

```bash
go run ./cmd/api
```

You should see:

Server running on :8080

## API Endpoints

### Users

Register: POST /users/register
JSON body:

```json
{
  "firstName": "Gizem",
  "lastName": "Güneş",
  "email": "gizem@example.com",
  "password": "123456"
}
```

Login: POST /users/login
JSON body:

```json
{
  "email": "gizem@example.com",
  "password": "123456"
}
```

### Products

Create Product: POST /products
JSON body:

```json
{
  "id": "uuid",
  "name": "Product Name",
  "description": "Description",
  "price": 99.99,
  "category": "Category"
}
```

List Products: GET /products

Update Product: PUT /products/{id}

Delete Product: DELETE /products/{id}

### Card

Add to Cart: POST /cart
JSON body:

```json
{
  "cartID": "user-id",
  "productID": "product-id",
  "quantity": 2,
  "unitPrice": 99.99
}
```
