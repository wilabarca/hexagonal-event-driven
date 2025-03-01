package application

import (
	entities "Event/src/Book/Domain/Entities"
	repositories "Event/src/Book/Domain/Repositories"
    "fmt"
)


type BookService struct {
	repository repositories.BookRepository
}


func NewBookService(repo repositories.BookRepository) *BookService {
	return &BookService{repository: repo}
}

// Crear un libro
func (b *BookService) CreateBook(book *entities.Book) error {
	return b.repository.CreateBook(book)
}

// Obtener un libro por ID
func (b *BookService) GetByID(id int64) (*entities.Book, error) {
	return b.repository.GetByID(id)
}

// Actualizar un libro
func (b *BookService) UpdateBook(book *entities.Book) error {
	if book.ID == 0 {
		return fmt.Errorf("book ID is required for update")
	}
	return b.repository.UpdateBook(book)
}


// Eliminar un libro
func (b *BookService) DeleteBook(id int64) error {
	return b.repository.DeleteBook(id)
}

// Obtener todos los libros
func (b *BookService) GetAll() ([]entities.Book, error) {
	return b.repository.GetAll()
}