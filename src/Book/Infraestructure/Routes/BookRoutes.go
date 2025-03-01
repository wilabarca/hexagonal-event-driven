package routes

import (
	controller "Event/src/Book/Infraestructure/Controller"
    "github.com/gin-gonic/gin"
)
func RegisterBookRoutes(router *gin.Engine, BookController *controller.BookController){
	routes := router.Group("Book")
	{
	   routes.GET("/", BookController.GetAllBooks)
	   routes.GET("/:id", BookController.GetBookByID)  
	   routes.POST("/", BookController.CreateBook)
	   routes.PUT("/:id", BookController.UpdateBook)
	   routes.DELETE("/:id", BookController.DeleteBook) 
	   
   }
}
