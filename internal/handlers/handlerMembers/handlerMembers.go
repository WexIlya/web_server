package handlerMembers

import (
	"net/http"
	storage "services/internal/storage/funcTables/funcMember"
)

func MembersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		storage.FindAllMembers()
	case http.MethodPost:
		storage.InsertMember(w, r)
	case http.MethodDelete:
		storage.DeleteMemberByID(w, r)
	}
}
