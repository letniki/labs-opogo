package rest

type DepositType struct {
	ID   int     `db:"id" json:"id"`
	Name string  `db:"name" json:"name"`
	Rate float64 `db:"rate" json:"rate"`
}

type Person struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	DepositID int    `db:"id" json:"deposit_id"`
}
