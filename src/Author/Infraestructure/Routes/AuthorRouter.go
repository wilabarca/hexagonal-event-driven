package routers

import (
	controller "Event/src/Author/Infraestructure/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterAuthorRoutes(router *gin.Engine, AuthorController *controller.AuthorController){
    AuthorGroup := router.Group("/Author")
    {
        AuthorGroup.GET("/", AuthorController.GetAllAuthors)
        AuthorGroup.GET("/:id", AuthorController.GetAuthorByID) 
        AuthorGroup.POST("/", AuthorController.CreateAuthor)
        AuthorGroup.PUT("/:id", AuthorController.UpdateAuthor)
        AuthorGroup.DELETE("/:id", AuthorController.DeleteAuthor)
    }
}
