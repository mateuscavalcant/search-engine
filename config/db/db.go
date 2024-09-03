package db

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func InitializeDB() {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DB")

	dbSource := user + ":" + password + "@tcp(localhost:3306)/" + database

	// Abre uma conexão com o banco de dados MySQL.
	conn, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Define o número máximo de conexões abertas.
	conn.SetMaxOpenConns(100)

	// Testa a conexão com o banco de dados para garantir que a conexão foi bem-sucedida.
	err = conn.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	db = conn
}

// GetDB retorna a conexão com o banco de dados MySQL.
func GetDB() *sql.DB {
	// Retorna a conexão existente com o banco de dados.
	return db
}
