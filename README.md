# E-Commerce System with Go

This is a simple e-commerce system built with **Go**, **PostgreSQL**, and **Docker**. It includes user registration, login, product management, and a shopping cart.

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

1. Pull the PostgreSQL image:

```bash
docker pull postgres
