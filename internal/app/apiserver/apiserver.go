package apiserver

import (
	"GoHttpRestApi/internal/app/store/sqlstore"
	"database/sql"
	"fmt"
	"net/http"
)

// Start ...
func Start(config *Config) error {
	db, err := newDb(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDb(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Hello db")
		return nil, err
	}

	return db, nil
}
