package main

import (
	"Event/core"
	AuthorService "Event/src/Author/Application"
	services "Event/src/Author/Application/Services"
	AuthorControlller "Event/src/Author/Infraestructure/Controller"
	AuthorDb "Event/src/Author/Infraestructure/Database"
	AuthorRoutes "Event/src/Author/Infraestructure/Routes"
	Repositories "Event/src/Author/Application/Repositories"
	BookService "Event/src/Book/Application"
	BookController "Event/src/Book/Infraestructure/Controller"
	BookDb "Event/src/Book/Infraestructure/Database"
	BookRoutes "Event/src/Book/Infraestructure/Routes"
    Adapters "Event/src/Author/Infraestructure/Adapters"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		return
	}
	rabbitMQURL := "amqp://guest:guest@52.55.188.45:5672/"
	rabbitAdapter, err := Adapters.NewRabbitMQAdapter(rabbitMQURL, "myQueue")
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ adapter: %v", err)
		return
	}
	rabbitRepo := Repositories.NewRabbitRepository(rabbitAdapter)  // Pasar el adaptador a RabbitRepository

	// Crear los servicios y controladores
	eventService := services.NewEventService(rabbitRepo)
	AuthorRepo := AuthorDb.NewsqlAuthorRepository(db)
	AuthorService := AuthorService.NewAuthorService(AuthorRepo)
	AuthorControlller := AuthorControlller.NewAuthorController(AuthorService, eventService)

	BookRepo := BookDb.NewsqlBookRepository(db)
	BookService := BookService.NewBookService(BookRepo)
	BookController := BookController.NewBookController(BookService)

	// Configurar el router y las rutas
	router := gin.Default()
	AuthorRoutes.RegisterAuthorRoutes(router, AuthorControlller)
	BookRoutes.RegisterBookRoutes(router, BookController)

	// Iniciar el servidor
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return
	}
}
