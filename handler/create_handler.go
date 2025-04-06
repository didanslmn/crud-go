package handler

import (
	"database/sql"
	"net/http"

	"github.com/didanslmn/crud-go/handler/template"
	"github.com/didanslmn/crud-go/model"
)

func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handleCreatePost(w, r, db)
			return
		}
		template.Render(w, "create.html", nil)
	}
}

func handleCreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	employee := model.Employee{
		Nama:    r.FormValue("nama"),
		NPWP:    r.FormValue("npwp"),
		Address: r.FormValue("address"),
	}
	if err := employee.Save(db); err != nil {
		http.Error(w, "Error creating data: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
