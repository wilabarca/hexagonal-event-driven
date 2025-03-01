package repositories

import entities "Event/src/Author/Domain/Entities"

type AuthorRepository interface {
	CreateAuthor(author *entities.Author) error
	GetAuthorByID(id int16) (*entities.Author, error)
	UpdateAuthor(author *entities.Author) error
	DeleteAuthor(id int16) error
	GetAllAuthor() ([]entities.Author, error)
}
