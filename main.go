package main

import (
	"net/http"

	"github.com/didanslmn/crud-go/database"
	"github.com/didanslmn/crud-go/routes"
)

func main() {
	db := database.InitDatabase()
	server := http.NewServeMux()
	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
