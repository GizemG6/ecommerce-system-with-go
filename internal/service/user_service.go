package service

import (
	"errors"

	"github.com/google/uuid"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"
	"github.com/GizemG6/ecommerce-system-with-go.git/internal/repository"
)

type UserService struct {
	repo repository.UserRepo
}

func NewUserService(r repository.UserRepo) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Register(firstName, lastName, email, password string) (*domain.User, error) {
	if firstName == "" || lastName == "" || email == "" || password == "" {
		return nil, errors.New("all fields are required")
	}

	existing, _ := s.repo.GetByEmail(email)
	if existing != nil {
		return nil, errors.New("email already exists")
	}

	u := &domain.User{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	if err := s.repo.Create(u); err != nil {
		return nil, err
	}

	return u, nil
}
func (s *UserService) Login(email, password string) (*domain.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email & password required")
	}

	u, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("user not found")
	}

	if u.Password != password {
		return nil, errors.New("invalid password")
	}

	return u, nil
}

func (s *UserService) ListUsers() ([]*domain.User, error) {
	users, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	result := make([]*domain.User, len(users))
	for i := range users {
		result[i] = &users[i]
	}
	return result, nil
}

func (s *UserService) UpdateUser(u *domain.User) error {
	return s.repo.Update(u)
}
