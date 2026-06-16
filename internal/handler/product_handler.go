package handler

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"go-crud-app/internal/model"
	"go-crud-app/internal/repository"
)

type ProductHandler struct {
	repo *repository.ProductRepository
	tmpl *template.Template
}

func NewProductHandler(repo *repository.ProductRepository, tmpl *template.Template) *ProductHandler {
	return &ProductHandler{repo: repo, tmpl: tmpl}
}

type PageData struct {
	Title    string
	Products []model.Product
	Product  *model.Product
	Error    string
	Success  string
}

// GET / — List all products
func (h *ProductHandler) Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	products, err := h.repo.GetAll()
	if err != nil {
		log.Println("Error fetching products:", err)
		products = []model.Product{}
	}

	data := PageData{Title: "Daftar Produk", Products: products}
	if err := h.tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GET /products/create — Show create form
func (h *ProductHandler) CreateForm(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Tambah Produk"}
	h.tmpl.ExecuteTemplate(w, "form.html", data)
}

// POST /products/create — Process create
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		data := PageData{Title: "Tambah Produk", Error: "Harga tidak valid"}
		h.tmpl.ExecuteTemplate(w, "form.html", data)
		return
	}

	stock, err := strconv.Atoi(r.FormValue("stock"))
	if err != nil {
		data := PageData{Title: "Tambah Produk", Error: "Stok tidak valid"}
		h.tmpl.ExecuteTemplate(w, "form.html", data)
		return
	}

	p := &model.Product{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Stock:       stock,
	}

	if p.Name == "" {
		data := PageData{Title: "Tambah Produk", Error: "Nama produk wajib diisi"}
		h.tmpl.ExecuteTemplate(w, "form.html", data)
		return
	}

	if err := h.repo.Create(p); err != nil {
		log.Println("Error creating product:", err)
		data := PageData{Title: "Tambah Produk", Error: "Gagal menyimpan produk"}
		h.tmpl.ExecuteTemplate(w, "form.html", data)
		return
	}

	http.Redirect(w, r, "/?success=Produk+berhasil+ditambahkan", http.StatusSeeOther)
}

// GET /products/edit?id=1 — Show edit form
func (h *ProductHandler) EditForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	p, err := h.repo.GetByID(id)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{Title: "Edit Produk", Product: p}
	h.tmpl.ExecuteTemplate(w, "form.html", data)
}

// POST /products/edit — Process edit
func (h *ProductHandler) EditProduct(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	stock, err := strconv.Atoi(r.FormValue("stock"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	p := &model.Product{
		ID:          id,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Stock:       stock,
	}

	if err := h.repo.Update(p); err != nil {
		log.Println("Error updating product:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/?success=Produk+berhasil+diperbarui", http.StatusSeeOther)
}

// POST /products/delete — Process delete
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		log.Println("Error deleting product:", err)
	}

	http.Redirect(w, r, "/?success=Produk+berhasil+dihapus", http.StatusSeeOther)
}
