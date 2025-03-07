package infraestructure

import (
	"api/src/Products/application"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ProductMessage representa el mensaje recibido en la cola
type ProductMessage struct {
	Id     int32   `json:"Id"`
	Name   string  `json:"Name"`
	Price  float32 `json:"Price"`
	Status string  `json:"Status"`
}

// RabbitMQConsumer estructura para consumir la cola
type RabbitMQConsumer struct {
	conn           *amqp.Connection
	channel        *amqp.Channel
	queueName      string
	createUseCase  *application.CreateProductUsecase
	getAllUseCase  *application.GetAllProduct
	getByIdUseCase *application.GetByIdProduct
	updateUseCase  *application.UpdateProduct
	deleteUseCase  *application.DeleteProductUsecase
}

// NewRabbitMQConsumer crea un nuevo consumidor de RabbitMQ
func NewRabbitMQConsumer(url, queueName string, createUseCase *application.CreateProductUsecase, getAllUseCase *application.GetAllProduct, getByIdUseCase *application.GetByIdProduct, updateUseCase *application.UpdateProduct, deleteUseCase *application.DeleteProductUsecase) (*RabbitMQConsumer, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQConsumer{
		conn:           conn,
		channel:        ch,
		queueName:      queueName,
		createUseCase:  createUseCase,
		getAllUseCase:  getAllUseCase,
		getByIdUseCase: getByIdUseCase,
		updateUseCase:  updateUseCase,
		deleteUseCase:  deleteUseCase,
	}, nil
}

// Start inicia la escucha de la cola
func (r *RabbitMQConsumer) Start() {
	msgs, err := r.channel.Consume(
		r.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("[RabbitMQ] Error al registrar consumidor: %s", err)
	}

	go func() {
		for d := range msgs {
			var msg ProductMessage
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("[RabbitMQ] Error al decodificar mensaje: %s", err)
				continue
			}

			switch msg.Status {
			case "post":
				log.Printf("[RabbitMQ] Guardando producto: %s", msg.Name)
				r.createUseCase.Execute(msg.Name, msg.Price)

			case "getById":
				log.Printf("[RabbitMQ] Buscando producto con ID: %d", msg.Id)
				product, err := r.getByIdUseCase.Execute(msg.Id)
				if err != nil {
					log.Printf("[RabbitMQ] Error al obtener producto: %s", err)
				} else {
					log.Printf("[RabbitMQ] Producto encontrado: %+v", product)
				}

			case "put":
				log.Printf("[RabbitMQ] Actualizando producto con ID: %d", msg.Id)
				err := r.updateUseCase.Execute(msg.Id, msg.Name, msg.Price)
				if err != nil {
					log.Printf("[RabbitMQ] Error al actualizar producto: %s", err)
				}

			case "delete":
				log.Printf("[RabbitMQ] Eliminando producto con ID: %d", msg.Id)
				err := r.deleteUseCase.Execute(msg.Id)
				if err != nil {
					log.Printf("[RabbitMQ] Error al eliminar producto: %s", err)
				}
			}
		}
	}()

	log.Println("[RabbitMQ] Escuchando mensajes...")
}

// Close cierra la conexión a RabbitMQ
func (r *RabbitMQConsumer) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
