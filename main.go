package main

import (
	"errors"
	"fmt"

	"github.com/guillermo-st/trip-agency-api/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			err := errors.New("Panicked while trying to initialize the WebAPI. Shutting down.")
			fmt.Println(err.Error(), r)
		}
		return
	}()

	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=local sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	srv, err := server.DefaultServer(db)
	if err != nil {
		panic(err)
	}

	err = srv.Run(8000)
	if err != nil {
		panic(err)
	}
}
