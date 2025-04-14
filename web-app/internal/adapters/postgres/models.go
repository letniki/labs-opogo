package postgres

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
