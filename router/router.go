package router

import (
	"database/sql"
	"net/http"

	"github.com/didanslmn/crud-go/handler"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(db *sql.DB) *chi.Mux {
	// routing
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/employees", http.StatusSeeOther)
	})

	r.Get("/employees", handler.IndexHandler(db))
	r.Get("/employees/create", handler.CreateHandler(db))
	r.Post("/employees/create", handler.CreateHandler(db))

	r.Get("/employees/edit/{id}", handler.EditHandler(db))
	r.Post("/employees/edit/{id}", handler.EditHandler(db))

	r.Post("/employees/delete/{id}", handler.DeleteHandler(db))

	return r
}
