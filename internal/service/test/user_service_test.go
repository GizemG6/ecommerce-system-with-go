package test

import (
	"errors"
	"testing"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"
	service "github.com/GizemG6/ecommerce-system-with-go.git/internal/service"
)

type MockUserRepo struct {
	CreateFn     func(u *domain.User) error
	GetByEmailFn func(email string) (*domain.User, error)
	ListFn       func() ([]domain.User, error)
	UpdateFn     func(u *domain.User) error
}

func (m *MockUserRepo) Create(u *domain.User) error {
	return m.CreateFn(u)
}

func (m *MockUserRepo) GetByEmail(email string) (*domain.User, error) {
	return m.GetByEmailFn(email)
}

func (m *MockUserRepo) List() ([]domain.User, error) {
	return m.ListFn()
}

func (m *MockUserRepo) Update(u *domain.User) error {
	return m.UpdateFn(u)
}

func TestRegister(t *testing.T) {
	mockRepo := &MockUserRepo{
		GetByEmailFn: func(email string) (*domain.User, error) {
			return nil, nil
		},
		CreateFn: func(u *domain.User) error { return nil },
	}

	s := service.NewUserService(mockRepo)

	u, err := s.Register("John", "Doe", "john@example.com", "123456")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.Email != "john@example.com" {
		t.Errorf("expected email 'john@example.com', got %s", u.Email)
	}
}

func TestLogin(t *testing.T) {
	mockRepo := &MockUserRepo{
		GetByEmailFn: func(email string) (*domain.User, error) {
			if email == "john@example.com" {
				return &domain.User{Email: email, Password: "123456"}, nil
			}
			return nil, errors.New("not found")
		},
	}

	s := service.NewUserService(mockRepo)

	u, err := s.Login("john@example.com", "123456")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.Email != "john@example.com" {
		t.Errorf("expected email 'john@example.com', got %s", u.Email)
	}
}

func TestListUsers(t *testing.T) {
	mockRepo := &MockUserRepo{
		ListFn: func() ([]domain.User, error) {
			return []domain.User{
				{Email: "u1@example.com"},
				{Email: "u2@example.com"},
			}, nil
		},
	}

	s := service.NewUserService(mockRepo)

	users, err := s.ListUsers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}
