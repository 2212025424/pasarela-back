package invoice

import (
	"context"
	"database/sql"
	"fmt"

	"pasarelapago.go/model"
	"pasarelapago.go/storage/postgres"
)

const (
	queryInsert  = "INSERT INTO invoice (id, invoice_date, email, is_product, is_subscription, product_id, subscription_id, price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	querySelect  = "SELECT id, invoice_date, email, is_product, is_subscription, product_id, subscription_id, price, created_at, updated_at FROM invoice"
	queryByID    = querySelect + " WHERE id = $1"
	queryByEmail = querySelect + " WHERE email = $1"
)

type Invoice struct {
	db *sql.DB
}

func New(db *sql.DB) Invoice {
	return Invoice{db: db}
}

func (i Invoice) Create(invoice *model.Invoice) error {
	emptyContext := context.Background()
	stmt, err := i.db.PrepareContext(emptyContext, queryInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.ExecContext(
		emptyContext,
		invoice.ID,
		invoice.InvoiceDate,
		invoice.Email,
		invoice.IsProduct,
		invoice.IsSubscription,
		invoice.ProductID,
		invoice.SubscriptionID,
		invoice.Price,
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

func (i Invoice) ByID(ID string) (model.Invoice, error) {
	emptyContext := context.Background()
	stmt, err := i.db.PrepareContext(emptyContext, queryByID)
	if err != nil {
		return model.Invoice{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(emptyContext, ID)
	if err != nil {
		return model.Invoice{}, err
	}

	return i.scan(row)
}

func (i Invoice) ByEmail(email string) (model.Invoices, error) {
	emptyContext := context.Background()
	stmt, err := i.db.PrepareContext(emptyContext, queryByEmail)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(emptyContext, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices model.Invoices
	for rows.Next() {
		invoice, err := i.scan(rows)
		if err != nil {
			return nil, err
		}

		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

func (i Invoice) scan(row postgres.RowScanner) (model.Invoice, error) {
	updatedAtNull := sql.NullTime{}
	invoice := model.Invoice{}

	err := row.Scan(
		&invoice.ID,
		&invoice.InvoiceDate,
		&invoice.Email,
		&invoice.IsProduct,
		&invoice.IsSubscription,
		&invoice.ProductID,
		&invoice.SubscriptionID,
		&invoice.Price,
		&invoice.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return model.Invoice{}, err
	}

	invoice.UpdatedAt = updatedAtNull.Time

	return invoice, nil
}
