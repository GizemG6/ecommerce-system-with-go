package postgres

import (
	"database/sql"
	"errors"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Create(u *domain.User) error {
	query := `
        INSERT INTO users (id, first_name, last_name, email, password, created_at)
        VALUES ($1, $2, $3, $4, $5, NOW())
    `
	_, err := r.db.Exec(query, u.ID, u.FirstName, u.LastName, u.Email, u.Password)
	return err
}

func (r *PostgresUserRepo) GetByEmail(email string) (*domain.User, error) {
	query := `
        SELECT id, first_name, last_name, email, password, created_at
        FROM users
        WHERE email = $1
    `
	row := r.db.QueryRow(query, email)

	u := &domain.User{}
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *PostgresUserRepo) List() ([]domain.User, error) {
	rows, err := r.db.Query(`SELECT id, first_name, last_name, email, password FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *PostgresUserRepo) Update(u *domain.User) error {
	if u.ID == "" {
		return errors.New("user id required")
	}

	query := `
        UPDATE users
        SET first_name=$1, last_name=$2, email=$3, password=$4
        WHERE id=$5
    `
	_, err := r.db.Exec(query, u.FirstName, u.LastName, u.Email, u.Password, u.ID)
	return err
}
