package repositories

import entities "Event/src/Book/Domain/Entities"


type BookRepository interface {
	CreateBook(book *entities.Book) error     
	GetAll() ([]entities.Book, error)         
	GetByID(id int64) (*entities.Book, error) 
	UpdateBook(book *entities.Book) error     
	DeleteBook(id int64) error                
}