package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/didanslmn/crud-go/model"
	"github.com/go-chi/chi/v5"
)

func DeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid employee id", http.StatusBadRequest)
			return
		}
		if err := model.Delete(db, id); err != nil {
			http.Error(w, "error deleting employe: "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/employees", http.StatusSeeOther)
	}
}
