package routes

import (
	"database/sql"
	"net/http"

	"github.com/GuidoWinata/crud-golang/controller"
)

func MapRouterp(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/" ,controller.HelloWorld())
	server.HandleFunc("/product" ,controller.IndexProduct(db))
	server.HandleFunc("/product/create" ,controller.ProductController(db))
	server.HandleFunc("/product/update" ,controller.UpdateProductController(db))
	server.HandleFunc("/product/delete" ,controller.DeleteProductController(db))
}