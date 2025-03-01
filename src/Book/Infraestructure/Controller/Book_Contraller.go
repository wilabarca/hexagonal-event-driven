package controller

import (
	application "Event/src/Book/Application"
	entities "Event/src/Book/Domain/Entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type BookController struct {
	service *application.BookService
}

func NewBookController(service *application.BookService) *BookController {
	return &BookController{service: service}
}

func (pc *BookController) CreateBook(c *gin.Context) {
	var book entities.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := pc.service.CreateBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book Created"})
}

func (pc *BookController) GetAllBooks(c *gin.Context) {
	books, err := pc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (pc *BookController) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book entities.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entry"})
		return
	}

	book.ID =int64(num) 

	err = pc.service.UpdateBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book Updated"})
}

func (pc *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = pc.service.DeleteBook(int64(num))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book Deleted"})
}

func (pc *BookController) GetBookByID(c *gin.Context) {
    id := c.Param("id") 

    num, err := strconv.Atoi(id) 
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    book, err := pc.service.GetByID(int64(num)) 
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if book == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    c.JSON(http.StatusOK, book) // Devolvemos el libro encontrado
}
