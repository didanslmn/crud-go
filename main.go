package main

import (
	"log"
	"net/http"

	"github.com/didanslmn/crud-go/database"
	"github.com/didanslmn/crud-go/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	// initial database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
