package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func ProductController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		r.ParseForm()
		nama := r.Form["nama"][0]
		harga := r.Form["harga"][0]
		deskripsi := r.Form["deskripsi"][0]

		_, err := db.Exec("INSERT INTO product (nama, harga, deskripsi) VALUES (?, ?, ?)", nama, harga, deskripsi)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/product", http.StatusMovedPermanently)
		return
	} else if r.Method == "GET" {
		fp := filepath.Join("views", "product.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	
		 err = tmpl.Execute(w, nil)
	
		 if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	
		}
	}

}