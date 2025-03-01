package core


import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error cargando archivo .env: %v", err)
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME")

	// Abrir la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al conectar la base de datos con DSN '%v': %v", dsn, err)
	}

	// Crear tablas si no existen
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("error creando tablas: %v", err)
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	authorTable := `CREATE TABLE IF NOT EXISTS authors (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL
	)`

	bookTable := `CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		year INT NOT NULL
	)`

	// Ejecutar creación de tabla authors
	if _, err := db.Exec(authorTable); err != nil {
		return fmt.Errorf("error creando tabla authors: %v", err)
	}

	// Ejecutar creación de tabla books
	if _, err := db.Exec(bookTable); err != nil {
		return fmt.Errorf("error creando tabla books: %v", err)
	}

	return nil
}