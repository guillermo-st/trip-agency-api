package main

import (
	"github.com/guillermo-st/trip-agency-api/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=local sslmode=disable")
	//db, err := sqlx.Open("postgres", "host=<yourhost> port=<yourport> user=<youruser> password=<yourpassword> dbname=<yourdbname> sslmode=disable")
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
