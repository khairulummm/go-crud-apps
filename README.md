# 🚀 Go CRUD App

Aplikasi web CRUD sederhana menggunakan **Go (Golang)** murni tanpa framework, dengan database **SQLite**.

## Fitur
- ✅ **Create** — Tambah produk baru
- ✅ **Read** — Lihat daftar semua produk
- ✅ **Update** — Edit data produk
- ✅ **Delete** — Hapus produk

## Teknologi
| | |
|---|---|
| **Bahasa** | Go 1.21 |
| **Database** | SQLite (via go-sqlite3) |
| **Template** | Go `html/template` |
| **Frontend** | HTML + CSS (tanpa framework) |

## Cara Menjalankan

### 1. Via GitHub Codespaces (Tanpa Install Apapun)
1. Klik tombol **"Code"** di GitHub → **"Codespaces"** → **"Create codespace on main"**
2. Tunggu environment siap (±1 menit)
3. Di terminal Codespaces, jalankan:
```bash
go mod tidy
go run ./cmd/main.go
```
4. Codespaces akan otomatis membuka browser ke port 8080 ✅

### 2. Di Komputer Lokal
```bash
# Clone repository
git clone <url-repo-kamu>
cd go-crud-app

# Download dependencies
go mod tidy

# Jalankan aplikasi
go run ./cmd/main.go

# Buka browser: http://localhost:8080
```

## Struktur Project
```
go-crud-app/
├── cmd/
│   └── main.go              # Entry point
├── internal/
│   ├── model/
│   │   └── product.go       # Struct data produk
│   ├── repository/
│   │   └── product_repository.go  # Operasi database
│   └── handler/
│       └── product_handler.go     # HTTP handler
├── templates/
│   ├── layout.html          # Base layout
│   ├── index.html           # Halaman daftar produk
│   └── form.html            # Form tambah/edit
├── .devcontainer/
│   └── devcontainer.json    # Konfigurasi Codespaces
├── go.mod
├── Makefile
└── README.md
```

## Belajar dari Project Ini

| File | Konsep Go yang Dipelajari |
|---|---|
| `main.go` | `http.ServeMux`, routing, template parsing |
| `product_repository.go` | `database/sql`, CRUD query, struct scanning |
| `product_handler.go` | HTTP handler, form parsing, redirect |
| `product.go` | Struct, field tags |
| `*.html` | Go template syntax (`{{range}}`, `{{if}}`, `{{define}}`) |
