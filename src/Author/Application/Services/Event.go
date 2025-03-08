package services

import (
	repositories "Event/src/Author/Application/Repositories"
	entities "Event/src/Author/Domain/Entities"
	"log"
)

type EventService struct {
	rabbitRepo repositories.RabbitRepository
}


func NewEventService(rabbitRepo repositories.RabbitRepository) *EventService {
	return &EventService{rabbitRepo: rabbitRepo}
}



func (s *EventService) AuthorUpdated(author *entities.Author) error {
    // el payload para el evento
    eventPayload := map[string]interface{}{
        "id":    author.ID,
        "name":  author.Name,
        "email": author.Email,
    }

    // se publica el evento
    err := s.rabbitRepo.PublishEvent("AuthorUpdated", eventPayload)
    if err != nil {
        log.Printf("Error al enviar el evento AuthorUpdated: %v", err)
        return err
    }
    return nil
}

