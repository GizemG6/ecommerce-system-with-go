package repository

import "github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"

type UserRepo interface {
	Create(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	List() ([]domain.User, error)
}

type ProductRepo interface {
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id string) error
	List() ([]domain.Product, error)
	GetByID(id string) (*domain.Product, error)
}

type CartRepo interface {
	AddItem(cartID string, item domain.CartItem) error
	GetCart(cartID string) (*domain.Cart, error)
	ClearCart(cartID string) error
}
