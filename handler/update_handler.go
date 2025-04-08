package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/didanslmn/crud-go/handler/template"
	"github.com/didanslmn/crud-go/model"
)

func EditHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleEditGet(w, r, db)
		case http.MethodPost:
			handleEditPost(w, r, db)
		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleEditGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// validasi id
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "invalid eployee id", http.StatusBadRequest)
		} else {
			log.Printf("error fetching employee: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
	// ambil data empolyee
	employee, err := model.GetByID(db, id)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	// rander tamplate
	template.Render(w, "edit.html", employee)
}

func handleEditPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// validasi id
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalis employee id", http.StatusBadRequest)
		return
	}
	employee := model.Employee{
		ID:      id,
		Nama:    r.FormValue("nama"),
		NPWP:    r.FormValue("npwp"),
		Address: r.FormValue("address"),
	}
	// validasi field required
	if employee.Nama == "" || employee.NPWP == "" || employee.Address == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	// update data
	if err := employee.Update(db); err != nil {
		//log.Printf("Update error: %v", err)
		http.Error(w, "Error updating employee: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// redirect index.html
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
