package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			gender VARCHAR(10),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create Table:", err)
	}
	log.Println("Tabel 'users' created successfully!")
}

func ConnectDatabase() {
	dsn := "root:@tcp(localhost:3306)/go_restapi_fiber"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db
	if err = DB.Ping(); err != nil {
		log.Fatal("failed to ping database: ", err)
	}
	log.Println("Database Connected!")

	createTable()
}
