package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/didanslmn/crud-go/handler/template"
	"github.com/didanslmn/crud-go/model"
)

func EditHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handleEditPost(w, r, db)
			return
		}
		handleEditGet(w, r, db)
	}
}

func handleEditGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid eployee id", http.StatusBadRequest)
		return
	}
	employee, err := model.GetByID(db, id)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	template.Render(w, "edit.html", employee)
}

func handleEditPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
	if err := employee.Update(db); err != nil {
		http.Error(w, "Error updating employee: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
