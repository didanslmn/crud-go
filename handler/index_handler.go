package handler

import (
	"database/sql"
	"net/http"

	"github.com/didanslmn/crud-go/handler/tamplate"
	"github.com/didanslmn/crud-go/model"
)

func IndexHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := model.GetAll(db)
		if err != nil {
			http.Error(w, "error fetching employees: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tamplate.Render(w, "index.html", employees)
	}
}
