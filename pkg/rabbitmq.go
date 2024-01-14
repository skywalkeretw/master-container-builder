// Package pkg contains the RabbitMQ related code
package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQData represents the configuration and connection to RabbitMQ
type RabbitMQData struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Conn     *amqp.Connection
	Ch       *amqp.Channel
}

// NewRabbitMQ initializes a new RabbitMQData instance with default values
func NewRabbitMQ() RabbitMQData {
	return RabbitMQData{
		UserName: GetEnvSting("RABBITMQ_USERNAME", "guest"),
		Password: GetEnvSting("RABBITMQ_PASSWORD", "guest"),
		Host:     GetEnvSting("RABBITMQ_HOST", "localhost"),
		Port:     GetEnvInt("RABBITMQ_PORT", 5672),
	}
}

// getUrl returns the RabbitMQ connection URL
func (r RabbitMQData) getUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", r.UserName, r.Password, r.Host, r.Port)
}

// Dial establishes a connection to RabbitMQ
func (r RabbitMQData) Dial() {
	var err error
	r.Conn, err = amqp.Dial(r.getUrl())
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer r.Conn.Close()

	r.Ch, err = r.Conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}
	defer r.Ch.Close()
}

// ReceiveMessages listens for messages on the specified queue
func (r RabbitMQData) ReceiveMessages(queueName string) {
	var err error

	err = r.Ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to set QoS", err)
	}

	msgs, err := r.Ch.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to register a consumer", err)
	}

	var forever chan struct{}

	go func() {
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			fmt.Println(d)
			fmt.Printf("Received a message: %s\n", string(d.Body))
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

// SendMessage sends a message to the "rpc_queue"
func (r RabbitMQData) SendMessage() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data := struct{ Name string }{"john"}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	err = r.Ch.PublishWithContext(ctx,
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: "corrId",
			ReplyTo:       "",
			Body:          jsonData,
		})
}
