package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/letniki/labs-opogo/internal"
)

type Client struct {
	db *sqlx.DB
}

func NewClient(db *sqlx.DB) Client {
	return Client{db: db}
}

// Створення депозиту
func (c Client) CreateDeposit(ctx context.Context, deposit internal.DepositType) (int, error) {
	query := `INSERT INTO deposits (id, name, rate) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := c.db.QueryRowContext(ctx, query, deposit.ID, deposit.Name, deposit.Rate).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("не вдалося створити депозит: %w", err)
	}
	return id, nil
}

// Отримання депозит за ID
func (c Client) GetDeposit(ctx context.Context, id int) (internal.DepositType, error) {
	query := `SELECT id, name, rate FROM deposits WHERE id = $1`
	var deposit internal.DepositType
	err := c.db.GetContext(ctx, &deposit, query, id)
	if err != nil {
		return internal.DepositType{}, fmt.Errorf("не вдалося отримати категорію: %w", err)
	}
	return deposit, nil
}

// Створення клієнту
func (c Client) CreatePerson(ctx context.Context, person internal.Person) (int, error) {
	query := `INSERT INTO persons (id, name, deposit_id) VALUES ($1,$2, $3) RETURNING id`
	var id int
	err := c.db.QueryRowContext(ctx, query, person.ID, person.Name, person.DepositID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("не вдалося створити клієнта: %w", err)
	}
	return id, nil
}

// Отримання клієнта за ID
func (c Client) GetPerson(ctx context.Context, id int) (internal.Person, error) {
	query := `SELECT id, name, deposit_id FROM persons WHERE id = $1`
	var person internal.Person
	err := c.db.GetContext(ctx, &person, query, id)
	if err != nil {
		return internal.Person{}, fmt.Errorf("не вдалося отримати клієнта: %w", err)
	}
	return person, nil
}
