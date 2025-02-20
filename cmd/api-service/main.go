package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	storage "services/internal"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "123"
	dbname   = "users"
)

func main() {
	dbpath := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	ctx := context.Background()
	db, err := sql.Open("postgres", dbpath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Successfully connected!\n")
		fmt.Fprintf(w, "DB: squad, starter packs\n")
		fmt.Fprintf(w, storage.FindAllMembers(db))
		fmt.Fprintf(w, storage.FindAllPack(db))
	})
	http.ListenAndServe(":8082", nil)
	//storage.InsertPack(db)
	//storage.InsertMember(db)
	//storage.FindAllPack(db)
	//storage.DeleteMemberByID(db, 1)
	//storage.DeletePackByID(db, 3)
}
