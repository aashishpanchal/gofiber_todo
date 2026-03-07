package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect() {
	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	pgdb, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pgdb.Close(context.Background())
}
