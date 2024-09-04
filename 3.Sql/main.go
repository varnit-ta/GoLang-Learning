package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	ID        int
	Name      string
	Price     float64
	Available bool
	CreatedAt string
}

func main() {
	fmt.Println("SQL Tutorial")

	connStr := "postgres://postgres:123@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(db)

	// Insert a product
	product := Product{Name: "Cake", Price: 20, Available: true}
	pk := insertProduct(db, product)
	fmt.Printf("Inserted product ID = %d\n", pk)

	// Query the inserted product by primary key
	queryProduct, err := queryProductByID(db, pk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Queried product: %+v\n", queryProduct)

	// Query all products
	products, err := queryAllProducts(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All products:")
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}
}

func createProductTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS products (	
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			price NUMERIC(6, 2) NOT NULL,
			available BOOLEAN,
			created_at TIMESTAMP DEFAULT NOW()
		);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product table created successfully")
}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO products (name, price, available)
		VALUES ($1, $2, $3) RETURNING id;
	`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}

	return pk
}

func queryProductByID(db *sql.DB, id int) (Product, error) {
	query := `SELECT id, name, price, available, created_at FROM products WHERE id = $1;`

	var product Product
	err := db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Available, &product.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, nil
		}
		return product, err
	}

	return product, nil
}

func queryAllProducts(db *sql.DB) ([]Product, error) {
	query := `SELECT id, name, price, available, created_at FROM products;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Available, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
