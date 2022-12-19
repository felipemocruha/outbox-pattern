package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // init pgx driver
	"github.com/jmoiron/sqlx"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres(uri string) (*Postgres, error) {
	db, err := otelsqlx.Connect("pgx", uri)
	if err != nil {
		return nil, fmt.Errorf("database connection: %w", err)
	}

	return &Postgres{DB: db}, nil
}

func (p *Postgres) Close() error {
	return p.DB.Close()
}
