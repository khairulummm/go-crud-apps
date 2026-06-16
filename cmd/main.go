package main

import (
	"html/template"
	"log"
	"net/http"

	"go-crud-app/internal/handler"
	"go-crud-app/internal/repository"
)

func main() {
	// Init database SQLite
	db, err := repository.InitDB("./products.db")
	if err != nil {
		log.Fatal("Gagal membuka database:", err)
	}
	defer db.Close()

	// Load templates
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"formatPrice": func(price float64) string {
			return "Rp " + formatNumber(price)
		},
	}).ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Gagal load templates:", err)
	}

	// Setup repository & handler
	repo := repository.NewProductRepository(db)
	h := handler.NewProductHandler(repo, tmpl)

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Index)
	mux.HandleFunc("/products/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateProduct(w, r)
		} else {
			h.CreateForm(w, r)
		}
	})
	mux.HandleFunc("/products/edit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.EditProduct(w, r)
		} else {
			h.EditForm(w, r)
		}
	})
	mux.HandleFunc("/products/delete", h.DeleteProduct)

	// Static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Server berjalan di http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func formatNumber(n float64) string {
	s := ""
	i := int(n)
	for i > 0 {
		rem := i % 1000
		i /= 1000
		if i > 0 {
			s = "." + zeroPad(rem, 3) + s
		} else {
			s = itoa(rem) + s
		}
	}
	if s == "" {
		s = "0"
	}
	return s
}

func zeroPad(n, width int) string {
	s := itoa(n)
	for len(s) < width {
		s = "0" + s
	}
	return s
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	return s
}
