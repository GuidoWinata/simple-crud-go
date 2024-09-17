package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Product struct {
	Id string
	Nama string
	Harga string
	Deskripsi string
}

func IndexProduct(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.Id,
			&product.Nama,
			&product.Harga,
			&product.Deskripsi,
		)
		if err != nil {
			w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
		}

		products = append(products, product)
	}

	fp := filepath.Join("views", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := make(map[string]any)
	data["products"] = products

	 err = tmpl.Execute(w, data)

	 if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	}
}