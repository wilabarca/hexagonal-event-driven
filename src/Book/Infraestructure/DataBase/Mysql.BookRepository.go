package database

import (
	entities "Event/src/Book/Domain/Entities"
	repositories "Event/src/Book/Domain/Repositories"
   "database/sql"
	"log"
	"fmt"
)

type MySQLBookRepository struct {
	db *sql.DB
}


func NewsqlBookRepository(db *sql.DB) repositories.BookRepository {
	return &MySQLBookRepository{db: db}
}

// Crear un libro en la base de datos
func (m *MySQLBookRepository) CreateBook(book *entities.Book) error {
	_, err := m.db.Exec("INSERT INTO books (title, year) VALUES (?, ?)", book.Title, book.Year)
	if err != nil {
		log.Println("Error al insertar el libro:", err)
		return err
	}
	return nil
}

// Eliminar un libro de la base de datos por su ID
func (m *MySQLBookRepository) DeleteBook(id int64) error {
	_, err := m.db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar el libro:", err)
	}
	return err
}

// Obtener todos los libros de la base de datos
func (m *MySQLBookRepository) GetAll() ([]entities.Book, error) {
	rows, err := m.db.Query("SELECT id, title, year FROM books")
	if err != nil {
		log.Println("Error al obtener los libros:", err)
		return nil, err
	}
	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Year); err != nil {
			log.Println("Error al escanear los libros:", err)
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// Obtener un libro por su ID
func (m *MySQLBookRepository) GetByID(id int64) (*entities.Book, error) {
	var book entities.Book
	err := m.db.QueryRow("SELECT id, title, year FROM books WHERE id = ?", id).
		Scan(&book.ID, &book.Title, &book.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No se encontró el libro con ID:", id)
			return nil, fmt.Errorf("book with id %d not found", id)
		}
		log.Println("Error al obtener el libro:", err)
		return nil, err
	}
	return &book, nil
}

// Actualizar la información de un libro en la base de datos
func (m *MySQLBookRepository) UpdateBook(book *entities.Book) error {
	log.Printf("Intentando actualizar libro con ID: %d\n", book.ID) 

	result, err := m.db.Exec("UPDATE books SET title = ?, year = ? WHERE id = ?", book.Title, book.Year, book.ID)
	if err != nil {
		log.Println("Error al actualizar el libro:", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("No se encontró ningún libro con ese ID para actualizar")
		return fmt.Errorf("book with id %d not found", book.ID)
	}

	return nil
}
