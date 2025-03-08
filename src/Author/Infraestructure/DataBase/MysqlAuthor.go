package database

import (
	entities "Event/src/Author/Domain/Entities"
	repositories "Event/src/Author/Domain/Repositories"
	"database/sql"
	"fmt"
)

type MySQLAuthorRepository struct {
	db *sql.DB
}

func NewsqlAuthorRepository(db *sql.DB) repositories.AuthorRepository {
	return &MySQLAuthorRepository{db: db}
}

// CreateAuthor guarda un nuevo autor en la base de datos.
func (m *MySQLAuthorRepository) CreateAuthor(author *entities.Author) error {
	_, err := m.db.Exec("INSERT INTO authors (name, email) VALUES (?, ?)", author.Name, author.Email)
	return err
}

// GetAuthorByID obtiene un autor por su ID.
func (m *MySQLAuthorRepository) GetAuthorByID(id int16) (*entities.Author, error) {
	var author entities.Author
	err := m.db.QueryRow("SELECT id, name, email FROM authors WHERE id = ?", id).
		Scan(&author.ID, &author.Name, &author.Email)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

// UpdateAuthor actualiza la información de un autor en la base de datos.
func (m *MySQLAuthorRepository) UpdateAuthor(author *entities.Author) error {

	result, err := m.db.Exec("UPDATE authors SET name = ?, email = ? WHERE id = ?", author.Name, author.Email, author.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar el autor: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ningún autor encontrado con el ID %d", author.ID)
	}

	return nil
}

// GetAllAuthor obtiene todos los autores desde la base de datos.
func (m *MySQLAuthorRepository) GetAllAuthor() ([]entities.Author, error) {
	rows, err := m.db.Query("SELECT id, name, email FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []entities.Author
	for rows.Next() {
		var author entities.Author
		if err := rows.Scan(&author.ID, &author.Name, &author.Email); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

// DeleteAuthor elimina un autor de la base de datos por su ID.
func (m *MySQLAuthorRepository) DeleteAuthor(id int16) error {
	_, err := m.db.Exec("DELETE FROM authors WHERE id = ?", id)
	return err
}
