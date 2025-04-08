package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/didanslmn/crud-go/handler/template"
	"github.com/didanslmn/crud-go/model"
)

func IndexHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ambil data dari database
		employees, err := model.GetAll(db)
		if err != nil {
			// log error di server
			log.Printf("Error fetching employees: %v", err)
			// pesan error ke clinet
			http.Error(w, "error fetching employees: "+err.Error(), http.StatusInternalServerError)
			return
		}
		//Render template dengan data
		template.Render(w, "index.html", employees)
	}
}
