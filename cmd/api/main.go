package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/handler"
	"github.com/GizemG6/ecommerce-system-with-go.git/internal/repository/postgres"
	"github.com/GizemG6/ecommerce-system-with-go.git/internal/service"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=123456 dbname=ecommerce sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := postgres.NewPostgresUserRepo(db)
	productRepo := postgres.NewPostgresProductRepo(db)
	cartRepo := postgres.NewPostgresCartRepo(db)

	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo)

	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)
	cartHandler := handler.NewCartHandler(cartService)

	r := mux.NewRouter()

	r.HandleFunc("/users/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/users/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/users", userHandler.ListUsers).Methods("GET")

	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products", productHandler.ListProducts).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	r.HandleFunc("/cart", cartHandler.AddToCart).Methods("POST")
	r.HandleFunc("/cart/{userID}", cartHandler.GetCart).Methods("GET")
	r.HandleFunc("/cart/{userID}/clear", cartHandler.ClearCart).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
