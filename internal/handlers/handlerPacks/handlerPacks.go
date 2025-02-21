package handlerPack

import (
	"net/http"
	storage "services/internal/storage/funcTables/funcPack"
)

func PacksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		storage.FindAllPack()
	case http.MethodPost:
		storage.InsertPack(w, r)
	case http.MethodDelete:
		storage.DeletePackByID(w, r)
	}
}
