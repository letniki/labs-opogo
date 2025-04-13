package term

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	connStr     = "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable"
	numDeposits = 20
	numPersons  = 10000
	csvFilePath = "clients.csv"
)

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	ctx := context.Background()

	// Створюємо депозит
	fmt.Println("Generating deposits...")
	depositIDs, err := generateDeposits(ctx, db)
	if err != nil {
		log.Fatal("Failed to generate deposits:", err)
	}

	// Створюємо клієнтів
	fmt.Println("Generating persons...")
	if err := generatePersons(ctx, db, depositIDs); err != nil {
		log.Fatal("Failed to generate persons:", err)
	}

	// Експортуємо у CSV
	fmt.Println("Exporting to CSV...")
	if err := exportToCSV(ctx, db); err != nil {
		log.Fatal("Failed to export to CSV:", err)
	}

	fmt.Println("Done!")
}

func generateDeposits(ctx context.Context, db *sql.DB) ([]int, error) {
	depositIDs := []int{}
	for i := 1; i <= numDeposits; i++ {
		name := fmt.Sprintf("Deposit_%d", i)
		rate := rand.Float64() * 20 // Випадковий ставка від 0 до 20
		var id int
		err := db.QueryRowContext(ctx, `INSERT INTO deposits (name, rate) VALUES ($1, $2) RETURNING id`, name, rate).Scan(&id)
		if err != nil {
			return nil, err
		}
		depositIDs = append(depositIDs, id)
	}
	return depositIDs, nil
}

func generatePersons(ctx context.Context, db *sql.DB, depositIDs []int) error {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= numPersons; i++ {
		name := fmt.Sprintf("Person_%d", i)

		_, err := db.ExecContext(ctx, `INSERT INTO persons (id, name, deposit_id) VALUES ($1, $2, $3)`, name, depositIDs)
		if err != nil {
			return err
		}
	}

	return nil
}

func exportToCSV(ctx context.Context, db *sql.DB) error {
	rows, err := db.QueryContext(ctx, `
		SELECT p.id, p.name, c.name AS depositId, c.rate
		FROM persons p
		JOIN deposits c ON p.deposit_id = c.id
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	file, err := os.Create(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Записуємо заголовки
	writer.Write([]string{"ID", "Person Name", "DepositId", "Rate"})

	// Записуємо дані
	for rows.Next() {
		var id int
		var personName string
		var depositId int
		var rate float64

		if err := rows.Scan(&id, &personName, &depositId, &rate); err != nil {
			return err
		}

		record := []string{
			strconv.Itoa(id),
			personName,
			strconv.Itoa(depositId),
			fmt.Sprintf("%.2f", rate),
		}
		writer.Write(record)
	}

	return nil
}
