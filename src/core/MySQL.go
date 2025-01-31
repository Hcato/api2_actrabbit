package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectToDB() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSchema := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbSchema)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Printf("Error al abrir la conexión: %v", err)
		return nil, fmt.Errorf("error al abrir la conexión: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Printf("Error al validar la conexión con la base de datos: %v", err)
		return nil, fmt.Errorf("error al validar la conexión: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10)
	db.SetConnMaxLifetime(5 * 60)

	fmt.Println("Conexión exitosa a la base de datos")
	return db, nil
}
