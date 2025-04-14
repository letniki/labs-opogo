package internal

import (
	"context"
)

type Storage interface {
	CreateDeposit(ctx context.Context, deposit DepositType) (int, error)
	GetDeposit(ctx context.Context, id int) (DepositType, error)
	CreatePerson(ctx context.Context, person Person) (int, error)
	GetPerson(ctx context.Context, id int) (Person, error)
}

type DepositType struct {
	ID   int     `db:"id"`
	Name string  `db:"name"`
	Rate float64 `db:"rate"`
}

type Person struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	DepositID int    `db:"deposit_id"`
}
