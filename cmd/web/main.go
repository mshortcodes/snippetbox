package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		logger.Error("DB_PASSWORD not set")
		os.Exit(1)
	}

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", fmt.Sprintf("web:%s@/snippetbox?parseTime=true", dbPassword), "MySQL data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", slog.String("addr", *addr))
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
