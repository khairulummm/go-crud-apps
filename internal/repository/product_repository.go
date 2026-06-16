package repository

import (
	"database/sql"
	"time"

	"go-crud-app/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS products (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		name        TEXT NOT NULL,
		description TEXT,
		price       REAL NOT NULL DEFAULT 0,
		stock       INTEGER NOT NULL DEFAULT 0,
		created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	rows, err := r.db.Query(`SELECT id, name, description, price, stock, created_at, updated_at FROM products ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	var p model.Product
	err := r.db.QueryRow(`SELECT id, name, description, price, stock, created_at, updated_at FROM products WHERE id = ?`, id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Create(p *model.Product) error {
	now := time.Now()
	result, err := r.db.Exec(
		`INSERT INTO products (name, description, price, stock, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`,
		p.Name, p.Description, p.Price, p.Stock, now, now,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)
	p.CreatedAt = now
	p.UpdatedAt = now
	return nil
}

func (r *ProductRepository) Update(p *model.Product) error {
	now := time.Now()
	_, err := r.db.Exec(
		`UPDATE products SET name=?, description=?, price=?, stock=?, updated_at=? WHERE id=?`,
		p.Name, p.Description, p.Price, p.Stock, now, p.ID,
	)
	if err != nil {
		return err
	}
	p.UpdatedAt = now
	return nil
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM products WHERE id=?`, id)
	return err
}
