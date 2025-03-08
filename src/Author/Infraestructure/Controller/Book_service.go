package controller

import (
	application "Event/src/Author/Application"
	services "Event/src/Author/Application/Services"
	entities "Event/src/Author/Domain/Entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type AuthorController struct {
	service *application.AuthorService
    eventService *services.EventService
}


func NewAuthorController(service *application.AuthorService, eventService *services.EventService) *AuthorController {
	return &AuthorController{service: service,
	eventService: eventService,}
}

// Crear un autor
func (c *AuthorController) CreateAuthor(ctx *gin.Context) {
	var author entities.Author
	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := c.service.CreateAuthor(&author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Author Created"})
}

// Obtener todos los autores
func (c *AuthorController) GetAllAuthors(ctx *gin.Context) {
	authors, err := c.service.GetAllAuthor()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, authors)
}

// Obtener un autor por ID
func (c *AuthorController) GetAuthorByID(ctx *gin.Context) {
    id := ctx.Param("id") 
    num, err := strconv.Atoi(id) 
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"}) 
        return
    }

    author, err := c.service.GetAuthorByID(int16(num)) 
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) 
        return
    }
    ctx.JSON(http.StatusOK, author)
}


// UpdateAuthor maneja la actualización de un autor y emite un evento.
func (c *AuthorController) UpdateAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var author entities.Author
	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	author.ID = authorID

	// Actualizamos el autor
	err = c.service.UpdateAuthor(&author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Emitimos el evento de que el autor ha sido actualizado
	err = c.eventService.AuthorUpdated(&author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al emitir el evento AuthorUpdated"})
		return
	}

	// Responder con éxito
	ctx.JSON(http.StatusOK, gin.H{"message": "Autor actualizado y evento emitido"})
}


// Eliminar un autor
func (c *AuthorController) DeleteAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
		return
	}

	err = c.service.DeleteAuthor(int16(num))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Author Deleted"})
}

