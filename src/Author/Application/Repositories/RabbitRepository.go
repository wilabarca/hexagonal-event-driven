package repositories

import adapters "Event/src/Author/Infraestructure/Adapters"

type RabbitRepository interface {
	PublishEvent(eventName string, payload interface{}) error
}

type rabbitRepo struct {
	adapter *adapters.RabbitMQAdapter
}

// NewRabbitRepository recibe un RabbitMQAdapter y lo utiliza en el repositorio
func NewRabbitRepository(adapter *adapters.RabbitMQAdapter) RabbitRepository {
	return &rabbitRepo{
		adapter: adapter,
	}
}

func (r *rabbitRepo) PublishEvent(eventName string, payload interface{}) error {
	// Usar el adaptador para publicar el evento
	return r.adapter.PublishEvent(eventName, payload)
}