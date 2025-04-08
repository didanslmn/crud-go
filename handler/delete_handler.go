package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/didanslmn/crud-go/model"
	"github.com/go-chi/chi/v5"
)

func DeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// validasi http method
		if r.Method != http.MethodPost { // biasanya delete menggunakan POST form dalam HTML
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// validasi id
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid employee id", http.StatusBadRequest)
			return
		}
		// eksekusi delete
		if err := model.Delete(db, id); err != nil {
			log.Printf("Delete error: %v", err) // Log error di server untuk dev

			switch {
			case err == sql.ErrNoRows:
				http.Error(w, "Employee not found", http.StatusNotFound)
			default:
				// pesan untuk client
				http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
			}
			return
		}
		// Redirect ke halaman list employee setelah berhasil delete
		http.Redirect(w, r, "/employees", http.StatusSeeOther)
	}
}
