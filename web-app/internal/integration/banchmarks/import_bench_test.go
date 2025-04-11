package banchmarks

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/letniki/labs-opogo/internal"
	"github.com/letniki/labs-opogo/internal/adapters/postgres"
	"github.com/letniki/labs-opogo/internal/ports/ftp"
	"log"
	"testing"
)

func BenchmarkProductsImport(b *testing.B) {
	dbExec, err := newDB()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	productsClient := postgres.NewClient(dbExec)
	productsService := internal.NewService(productsClient)
	productsParser := ftp.NewParser(productsService)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err = productsParser.Run(ctx, `C:\Users\User\IdeaProjects\labs-opogo1\web-app\internal\integration\data\clients.csv`); err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func newDB() (*sqlx.DB, error) {
	dsn := "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable&search_path=persons"
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
