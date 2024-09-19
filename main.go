package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
		return
	}

	addr := os.Getenv("ADDR")
	dsn := os.Getenv("DATABASE_URL")
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloWorld)

	_, err = connectToDatabase(dsn)
	if err != nil {
		slog.Error("Unable to connect to the database", "err", err)
		return
	}

	slog.Info("Starting web server", "addr", addr)
	err = http.ListenAndServe(addr, mux)
	panic(err)
}

func connectToDatabase(dsn string) (*sql.DB, error) {
	if dsn == "" {
		return nil, errors.New("data source name is empty")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
