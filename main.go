package main

import (
	"log"
	"net/http"

	"github.com/didanslmn/crud-go/database"
	"github.com/didanslmn/crud-go/handler"
)

func main() {
	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/employees", http.StatusSeeOther)
	})

	http.HandleFunc("/employees", handler.IndexHandler(db))

	// Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
