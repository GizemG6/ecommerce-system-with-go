package test

import (
	"errors"
	"testing"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"
	service "github.com/GizemG6/ecommerce-system-with-go.git/internal/service"
)

// Mock repository for Product
type MockProductRepo struct {
	CreateFn  func(p *domain.Product) error
	GetByIDFn func(id string) (*domain.Product, error)
	ListFn    func() ([]domain.Product, error)
	UpdateFn  func(p *domain.Product) error
	DeleteFn  func(id string) error
}

func (m *MockProductRepo) Create(p *domain.Product) error {
	return m.CreateFn(p)
}

func (m *MockProductRepo) GetByID(id string) (*domain.Product, error) {
	return m.GetByIDFn(id)
}

func (m *MockProductRepo) List() ([]domain.Product, error) {
	return m.ListFn()
}

func (m *MockProductRepo) Update(p *domain.Product) error {
	return m.UpdateFn(p)
}

func (m *MockProductRepo) Delete(id string) error {
	return m.DeleteFn(id)
}

func TestCreateProduct(t *testing.T) {
	mockRepo := &MockProductRepo{
		CreateFn: func(p *domain.Product) error { return nil },
	}

	s := service.NewProductService(mockRepo)

	p, err := s.CreateProduct("Test Product", "Description", 10.0, "Category")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if p.Name != "Test Product" {
		t.Errorf("expected name 'Test Product', got %s", p.Name)
	}
}

func TestGetProduct(t *testing.T) {
	mockRepo := &MockProductRepo{
		GetByIDFn: func(id string) (*domain.Product, error) {
			if id == "123" {
				return &domain.Product{ID: "123", Name: "Test"}, nil
			}
			return nil, errors.New("not found")
		},
	}

	s := service.NewProductService(mockRepo)

	p, err := s.GetProduct("123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if p.ID != "123" {
		t.Errorf("expected ID '123', got %s", p.ID)
	}
}

func TestListProducts(t *testing.T) {
	mockRepo := &MockProductRepo{
		ListFn: func() ([]domain.Product, error) {
			return []domain.Product{
				{ID: "1", Name: "P1"},
				{ID: "2", Name: "P2"},
			}, nil
		},
	}

	s := service.NewProductService(mockRepo)

	products, err := s.ListProducts()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(products) != 2 {
		t.Errorf("expected 2 products, got %d", len(products))
	}
}
