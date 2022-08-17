package main

import (
	"errors"
	"fmt"

	"github.com/guillermo-st/trip-agency-api/database"
	"github.com/guillermo-st/trip-agency-api/server"
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

	db, err := database.NewPostgresDBClient()
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
