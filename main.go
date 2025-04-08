package main

import (
	"log"
	"net/http"

	"github.com/didanslmn/crud-go/database"
	"github.com/didanslmn/crud-go/router"
)

func main() {
	// initial database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// setup router
	r := router.SetupRouter(db)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
