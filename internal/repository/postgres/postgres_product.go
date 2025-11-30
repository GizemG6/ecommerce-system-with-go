package postgres

import (
	"database/sql"
	"errors"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"
)

type PostgresProductRepo struct {
	db *sql.DB
}

func NewPostgresProductRepo(db *sql.DB) *PostgresProductRepo {
	return &PostgresProductRepo{db: db}
}

func (r *PostgresProductRepo) Create(p *domain.Product) error {
	query := `
        INSERT INTO products (id, name, description, price, category)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := r.db.Exec(query, p.ID, p.Name, p.Description, p.Price, p.Category)
	return err
}

func (r *PostgresProductRepo) GetByID(id string) (*domain.Product, error) {
	query := `
        SELECT id, name, description, price, category
        FROM products
        WHERE id = $1
    `
	row := r.db.QueryRow(query, id)

	p := &domain.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *PostgresProductRepo) List() ([]domain.Product, error) {
	rows, err := r.db.Query(`
        SELECT id, name, description, price, category
        FROM products
    `)
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *PostgresProductRepo) Update(p *domain.Product) error {
	if p.ID == "" {
		return errors.New("product id required")
	}

	query := `
        UPDATE products
        SET name=$1, description=$2, price=$3, category=$4
        WHERE id=$5
    `
	_, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.Category, p.ID)
	return err
}

func (r *PostgresProductRepo) Delete(id string) error {
	if id == "" {
		return errors.New("id required")
	}

	_, err := r.db.Exec(`DELETE FROM products WHERE id = $1`, id)
	return err
}
