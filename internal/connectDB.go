package connect

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "*"
	port     = *
	user     = "*"
	password = "*"
	dbname   = "*"
)

var Connectdb *sql.DB

func ConnectDB() {
	dbpath := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	ctx := context.Background()
	db, err := sql.Open("postgres", dbpath)
	Connectdb = db

	if err != nil {
		log.Fatal(err)
	}
	if err := Connectdb.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
}
