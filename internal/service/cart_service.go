package service

import (
	"errors"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"
	"github.com/GizemG6/ecommerce-system-with-go.git/internal/repository"
)

type CartService struct {
	repo repository.CartRepo
}

func NewCartService(r repository.CartRepo) *CartService {
	return &CartService{repo: r}
}

func (s *CartService) AddToCart(cartID string, item domain.CartItem) error {
	if item.Quantity <= 0 {
		return errors.New("quantity must be > 0")
	}
	return s.repo.AddItem(cartID, item)
}

func (s *CartService) GetCart(cartID string) (*domain.Cart, error) {
	return s.repo.GetCart(cartID)
}

func (s *CartService) ClearCart(cartID string) error {
	return s.repo.ClearCart(cartID)
}
