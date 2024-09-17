package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func UpdateProductController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.URL.Query().Get("id")

		r.ParseForm()
		nama := r.Form["nama"][0]
		harga := r.Form["harga"][0]
		deskripsi := r.Form["deskripsi"][0]

		_, err := db.Exec("UPDATE product SET nama=?, harga=?, deskripsi=? WHERE id=?", nama, harga, deskripsi, id)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/product", http.StatusMovedPermanently)
		return
	} else if r.Method == "GET" {
		id := r.URL.Query().Get("id")

		row := db.QueryRow("SELECT * FROM product WHERE id=?", id)
		if row.Err() != nil {
			w.Write([]byte(row.Err().Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var product Product
		err := row.Scan(
			&product.Id,
			&product.Nama,
			&product.Harga,
			&product.Deskripsi,
		)
		product.Id = id
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fp := filepath.Join("views", "update.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["product"] = product
	
		 err = tmpl.Execute(w, data)
	
		 if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	
		}
	}

}