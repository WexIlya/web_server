package main

import (
	"fmt"
	"net/http"
	connect "services/internal"
	"services/internal/handlers/handlerMembers"
	handlerPack "services/internal/handlers/handlerPacks"
	storageM "services/internal/storage/funcTables/funcMember"
	storageP "services/internal/storage/funcTables/funcPack"

	_ "github.com/lib/pq"
)

func main() {
	connect.ConnectDB()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DB: squad\n")
		fmt.Fprintf(w, storageM.FindAllMembers())
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "DB: starter packs\n")
		fmt.Fprintf(w, storageP.FindAllPack())
		fmt.Fprintf(w, "\n")
	})
	http.HandleFunc("/squad", handlerMembers.MembersHandler)
	http.HandleFunc("/starterPacks", handlerPack.PacksHandler)
	http.ListenAndServe(":8082", nil)
}
