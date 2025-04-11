package main

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/letniki/labs-opogo/internal"
	"github.com/letniki/labs-opogo/internal/adapters/postgres"
	"github.com/letniki/labs-opogo/internal/ports/rest"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	app := newApplication()
	app.start()
}

type application struct {
	server *http.Server
}

func newApplication() application {
	db, err := newDB()
	if err != nil {
		log.Fatal(err)
	}

	server := newServer(db)

	return application{
		server: server,
	}
}

func newDB() (*sqlx.DB, error) {
	dsn := "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable&search_path=persons"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func newServer(db *sqlx.DB) *http.Server {
	client := postgres.NewClient(db)
	service := internal.NewService(client)
	handler := rest.NewHandler(service)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/deposits", handler.CreateDeposit).Methods("POST")
	r.HandleFunc("/api/v1/deposits/{id:[0-9]+}", handler.GetDeposit).Methods("GET")

	r.HandleFunc("/api/v1/persons", handler.CreatePerson).Methods("POST")
	r.HandleFunc("/api/v1/persons/{id:[0-9]+}", handler.GetPerson).Methods("GET")

	return &http.Server{
		Addr:    "localhost:8081",
		Handler: r,
	}
}

func (app application) start() {
	log.Println("Сервер запущено на http://localhost:8081")
	if err := app.server.ListenAndServe(); err != nil {
		log.Fatal("Помилка запуску сервера:", err)
	}
}
