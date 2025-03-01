package main

import (
	"Event/core"
	AuthorService "Event/src/Author/Application"
	AuthorControlller "Event/src/Author/Infraestructure/Controller"
	AuthorDb "Event/src/Author/Infraestructure/Database"
	AuthorRoutes "Event/src/Author/Infraestructure/Routes"

	BookService "Event/src/Book/Application"
	BookController "Event/src/Book/Infraestructure/Controller"
	BookDb "Event/src/Book/Infraestructure/Database"
	BookRoutes "Event/src/Book/Infraestructure/Routes"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect")
	    return
	}

	

	AuthorRepo :=  AuthorDb.NewsqlAuthorRepository(db)
	AuthorService :=  AuthorService.NewAuthorService(AuthorRepo)
	AuthorControlller := AuthorControlller.NewAuthorController(AuthorService)

	BookRepo := BookDb.NewsqlBookRepository(db)
	BookService := BookService.NewBookService(BookRepo)
	BookController := BookController.NewBookController(BookService)


	router := gin.Default()


	AuthorRoutes.RegisterAuthorRoutes(router , AuthorControlller)
	BookRoutes.RegisterBookRoutes(router, BookController)


	
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
        println(err)
		
	}

	
}