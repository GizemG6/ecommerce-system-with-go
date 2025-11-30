package service

import (
	"errors"

	"github.com/google/uuid"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"
	"github.com/GizemG6/ecommerce-system-with-go.git/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepo
}

func NewProductService(r repository.ProductRepo) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) CreateProduct(name, desc string, price float64, category string) (*domain.Product, error) {
	if name == "" || price <= 0 {
		return nil, errors.New("name and price are required")
	}

	p := &domain.Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: desc,
		Price:       price,
		Category:    category,
	}

	err := s.repo.Create(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *ProductService) UpdateProduct(p *domain.Product) error {
	if p.ID == "" {
		return errors.New("product id required")
	}
	return s.repo.Update(p)
}

func (s *ProductService) DeleteProduct(id string) error {
	if id == "" {
		return errors.New("id required")
	}
	return s.repo.Delete(id)
}

func (s *ProductService) GetProduct(id string) (*domain.Product, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.repo.GetByID(id)
}

func (s *ProductService) ListProducts() ([]*domain.Product, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Product, len(products))
	for i := range products {
		result[i] = &products[i]
	}
	return result, nil
}
