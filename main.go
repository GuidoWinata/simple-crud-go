package main

import (
	"net/http"

	"github.com/GuidoWinata/crud-golang/database"
	"github.com/GuidoWinata/crud-golang/routes"
)

func main() {
 	db := database.InitDatabase()
	server := http.NewServeMux()

	routes.MapRouterp(server, db)

	http.ListenAndServe(":8080", server)
}