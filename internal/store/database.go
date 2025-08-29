package store

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")

	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	fmt.Print("Connected to the db, yeay")
	return db, nil
}
