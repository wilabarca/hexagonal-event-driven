package application

import (
	entities "Event/src/Author/Domain/Entities"
	repositories "Event/src/Author/Domain/Repositories"
)

type AuthorService struct {
    repository repositories.AuthorRepository
}

func NewAuthorService(repo repositories.AuthorRepository) *AuthorService{
	return &AuthorService{repository: repo}
}

func (s *AuthorService) CreateAuthor(author *entities.Author) error{
    return s.repository.CreateAuthor(author)
}
func (s*AuthorService) GetAuthorByID(id int16) (*entities.Author, error){
	return s.repository.GetAuthorByID(id)
}

func (s *AuthorService) UpdateAuthor(author *entities.Author) error {
    err := s.repository.UpdateAuthor(author)
    if err != nil {
        return err 
    }
    return nil
}

func (s *AuthorService) DeleteAuthor (id int16) error{
	return s.repository.DeleteAuthor(id)
}

func (s *AuthorService) GetAllAuthor() ([]entities.Author, error){
	return s.repository.GetAllAuthor()
}


