package test

import (
	"errors"
	"testing"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"
	service "github.com/GizemG6/ecommerce-system-with-go.git/internal/service"
)

type MockCartRepo struct {
	carts map[string]*domain.Cart
}

func NewMockCartRepo() *MockCartRepo {
	return &MockCartRepo{
		carts: make(map[string]*domain.Cart),
	}
}

func (m *MockCartRepo) AddItem(cartID string, item domain.CartItem) error {
	if item.Quantity <= 0 {
		return errors.New("quantity must be > 0")
	}
	cart, exists := m.carts[cartID]
	if !exists {
		cart = &domain.Cart{ID: cartID, Items: []domain.CartItem{}}
		m.carts[cartID] = cart
	}
	cart.Items = append(cart.Items, item)
	return nil
}

func (m *MockCartRepo) GetCart(cartID string) (*domain.Cart, error) {
	cart, exists := m.carts[cartID]
	if !exists {
		return &domain.Cart{ID: cartID, Items: []domain.CartItem{}}, nil
	}
	return cart, nil
}

func (m *MockCartRepo) ClearCart(cartID string) error {
	delete(m.carts, cartID)
	return nil
}

func TestCartService_AddToCartAndGetCart(t *testing.T) {
	mockRepo := NewMockCartRepo()
	service := service.NewCartService(mockRepo)

	item := domain.CartItem{ProductID: "p1", Quantity: 2, UnitPrice: 10.0}
	err := service.AddToCart("cart1", item)
	if err != nil {
		t.Fatal(err)
	}

	cart, err := service.GetCart("cart1")
	if err != nil {
		t.Fatal(err)
	}

	if len(cart.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(cart.Items))
	}

	if cart.Items[0].Quantity != 2 {
		t.Fatalf("expected quantity 2, got %d", cart.Items[0].Quantity)
	}
}

func TestCartService_ClearCart(t *testing.T) {
	mockRepo := NewMockCartRepo()
	service := service.NewCartService(mockRepo)

	item := domain.CartItem{ProductID: "p1", Quantity: 1, UnitPrice: 5.0}
	service.AddToCart("cart1", item)

	err := service.ClearCart("cart1")
	if err != nil {
		t.Fatal(err)
	}

	cart, err := service.GetCart("cart1")
	if err != nil {
		t.Fatal(err)
	}

	if len(cart.Items) != 0 {
		t.Fatalf("expected 0 items after clear, got %d", len(cart.Items))
	}
}

func TestCartService_AddToCart_InvalidQuantity(t *testing.T) {
	mockRepo := NewMockCartRepo()
	service := service.NewCartService(mockRepo)

	item := domain.CartItem{ProductID: "p1", Quantity: 0, UnitPrice: 10.0}
	err := service.AddToCart("cart1", item)
	if err == nil {
		t.Fatal("expected error for quantity <= 0, got nil")
	}
}
