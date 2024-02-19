package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

type Postgres interface {
	Connect() error
	Ping() error
	Disconnect() error
}

type postgres struct {
	db *sql.DB
}

func NewPostgres() Postgres {
	return &postgres{}
}

var POSTGRES_DB *sql.DB

func (p *postgres) Connect() error {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	slog.Info("Connecting to Postgres", "host", host, "port", port, "user", user, "dbname", dbname)

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	slog.Info("Connected to Postgres")

	p.db = db
	POSTGRES_DB = db
	return nil
}

func (p *postgres) Ping() error {
	err := p.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (p *postgres) Disconnect() error {
	err := p.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetPostgres() *sql.DB {
	return POSTGRES_DB
}
