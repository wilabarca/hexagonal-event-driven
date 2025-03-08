package adapters

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// RabbitMQAdapter estructura para manejar RabbitMQ
type RabbitMQAdapter struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

// NewRabbitMQAdapter crea una nueva instancia del adaptador RabbitMQ
func NewRabbitMQAdapter(url, queueName string) (*RabbitMQAdapter, error) {
	// Intentar establecer la conexión con RabbitMQ
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	// Asegurarse de cerrar la conexión correctamente
	defer conn.Close()

	// Crear un canal de comunicación
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	// Asegurarse de cerrar el canal correctamente
	defer ch.Close()

	// Declarar la cola en RabbitMQ
	queue, err := ch.QueueDeclare(
		queueName, // nombre de la cola
		false,     // durable (la cola no sobrevive a reinicios de RabbitMQ)
		false,     // delete when unused (se eliminará cuando no esté en uso)
		false,     // exclusive (no se permite que otros clientes usen la cola)
		false,     // no-wait (no espera que la declaración sea confirmada)
		nil,       // argumentos
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %v", err)
	}

	// Devolver el adaptador creado
	return &RabbitMQAdapter{
		connection: conn,
		channel:    ch,
		queue:      queue,
	}, nil
}

// PublishEvent publica un evento en RabbitMQ
func (r *RabbitMQAdapter) PublishEvent(eventName string, payload interface{}) error {
	// Convertir los datos del evento a JSON
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal event data: %v", err)
	}

	// Publicar el evento en RabbitMQ
	err = r.channel.Publish(
		"",             // exchange (vacío por defecto)
		r.queue.Name,   // routing key (nombre de la cola)
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json", // tipo de contenido
			Body:        body,                // cuerpo del mensaje
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	log.Printf("Event %s sent successfully", eventName)
	return nil
}

// Close cierra la conexión y el canal de RabbitMQ
func (r *RabbitMQAdapter) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.connection != nil {
		r.connection.Close()
	}
}
