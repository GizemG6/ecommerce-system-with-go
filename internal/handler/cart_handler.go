package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"
	"github.com/GizemG6/ecommerce-system-with-go.git/internal/service"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(s *service.CartService) *CartHandler {
	return &CartHandler{service: s}
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var item struct {
		UserID    string  `json:"userID"`
		ProductID string  `json:"productID"`
		Quantity  int     `json:"quantity"`
		UnitPrice float64 `json:"unitPrice"`
	}

	json.NewDecoder(r.Body).Decode(&item)

	err := h.service.AddToCart(item.UserID, domain.CartItem{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		UnitPrice: item.UnitPrice,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Item added"))
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]
	cart, err := h.service.GetCart(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]
	err := h.service.ClearCart(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte("Cart cleared"))
}
