package handler

import (
	"database/sql"
	"net/http"

	"github.com/didanslmn/crud-go/handler/template"
	"github.com/didanslmn/crud-go/model"
)

func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Cek metod request
		if r.Method == http.MethodPost {
			handleCreatePost(w, r, db) // proses POST
			return
		}
		// Render template form create untuk GET request
		template.Render(w, "create.html", nil)
	}
}

func handleCreatePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Membuat objek employee dari form values
	employee := model.Employee{
		Nama:    r.FormValue("nama"),
		NPWP:    r.FormValue("npwp"),
		Address: r.FormValue("address"),
	}
	// Simpan employee ke database
	if err := employee.Save(db); err != nil {
		http.Error(w, "Error creating data: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Redirect ke halaman list employee
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
