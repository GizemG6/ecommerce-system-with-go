package postgres

import (
	"database/sql"

	"github.com/GizemG6/ecommerce-system-with-go.git/internal/domain"
)

type PostgresCartRepo struct {
	db *sql.DB
}

func NewPostgresCartRepo(db *sql.DB) *PostgresCartRepo {
	return &PostgresCartRepo{db: db}
}

func (r *PostgresCartRepo) AddItem(cartID string, item domain.CartItem) error {
	// create cart if not exists
	_, _ = r.db.Exec(`INSERT INTO carts (id) VALUES ($1) ON CONFLICT DO NOTHING`, cartID)

	query := `
        INSERT INTO cart_items (cart_id, product_id, quantity, unit_price)
        VALUES ($1, $2, $3, $4)
    `
	_, err := r.db.Exec(query, cartID, item.ProductID, item.Quantity, item.UnitPrice)
	return err
}

func (r *PostgresCartRepo) GetCart(cartID string) (*domain.Cart, error) {
	cart := &domain.Cart{
		ID:    cartID,
		Items: []domain.CartItem{},
	}

	row := r.db.QueryRow(`SELECT id FROM carts WHERE id=$1`, cartID)
	var tmp string
	err := row.Scan(&tmp)
	if err == sql.ErrNoRows {
		return cart, nil
	}
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(`
        SELECT product_id, quantity, unit_price
        FROM cart_items
        WHERE cart_id = $1
    `, cartID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var it domain.CartItem
		if err := rows.Scan(&it.ProductID, &it.Quantity, &it.UnitPrice); err != nil {
			return nil, err
		}
		cart.Items = append(cart.Items, it)
	}

	return cart, nil
}

func (r *PostgresCartRepo) ClearCart(cartID string) error {
	_, err := r.db.Exec(`DELETE FROM cart_items WHERE cart_id = $1`, cartID)
	return err
}
