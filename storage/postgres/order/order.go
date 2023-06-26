package order

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"pasarelapago.go/model"
	"pasarelapago.go/storage/postgres"
)

const (
	queryInsert = "INSERT INTO corder (id, email, is_product, is_subscription, product_id, type_subs, price) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	queryByID   = "SELECT id, email, is_product, is_subscription, product_id, type_subs, price, created_at, updated_at FROM corder WHERE id = $1"
)

type Order struct {
	db *sql.DB
}

func New(db *sql.DB) Order {
	return Order{db: db}
}

func (o Order) Create(order *model.Order) error {
	emptyContext := context.Background()
	stmt, err := o.db.PrepareContext(emptyContext, queryInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.ExecContext(
		emptyContext,
		order.ID,
		order.Email,
		order.IsProduct,
		order.IsSubscription,
		order.ProductID,
		order.TypeSubs,
		order.Price,
	)
	if err != nil {
		return err
	}

	got, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if got != 1 {
		return fmt.Errorf("expected 1 row, got %d", got)
	}

	return nil
}

func (o Order) ByID(ID uuid.UUID) (model.Order, error) {
	emptyContext := context.Background()
	stmt, err := o.db.PrepareContext(emptyContext, queryByID)
	if err != nil {
		return model.Order{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(emptyContext, ID)

	return o.scan(row)
}

func (o Order) scan(row postgres.RowScanner) (model.Order, error) {
	productIDNull := sql.NullString{}
	typeSubsNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	order := model.Order{}

	err := row.Scan(
		&order.ID,
		&order.Email,
		&order.IsProduct,
		&order.IsSubscription,
		&productIDNull,
		&typeSubsNull,
		&order.Price,
		&order.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return model.Order{}, err
	}

	order.ProductID = uuid.MustParse(productIDNull.String)
	order.TypeSubs = typeSubsNull.String

	return order, nil
}
