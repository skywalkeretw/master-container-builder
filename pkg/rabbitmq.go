package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type FuncData struct {
	Name         string
	Code         string
	Language     string
	OpenAPISpec  string
	AsyncAPISpec string
}

type RabbitMQData struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type ReturnInfo struct {
	ImageName string `json:"imagename"`
}

func newRabbitMQ() RabbitMQData {
	return RabbitMQData{
		UserName: GetEnvSting("RABBITMQ_USERNAME", "guest"),
		Password: GetEnvSting("RABBITMQ_PASSWORD", "guest"),
		Host:     GetEnvSting("RABBITMQ_HOST", "localhost"),
		Port:     GetEnvInt("RABBITMQ_PORT", 5672),
	}
}
func (r RabbitMQData) getUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", r.UserName, r.Password, r.Host, r.Port)

}

func ListenToQueue(queue string) {
	rmq := newRabbitMQ()
	conn, err := amqp.Dial(rmq.getUrl())
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to declare a queue", err)
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to set QoS", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to register a consumer", err)
	}

	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			fmt.Println(string(d.Body))
			var funcData FuncData
			err := json.Unmarshal(d.Body, &funcData)
			if err != nil {
				log.Panicf("%s: %s", "Failed to unmarshal Body", err)
			}

			newImage := NewPodmanImage(funcData)
			err = newImage.build()
			if err != nil {
				fmt.Printf("Failed to build %v", err)
			}
			imageName, err := newImage.push()
			if err != nil {
				fmt.Printf("Failed to push %v", err)
			}
			err = newImage.remove()
			if err != nil {
				fmt.Printf("Failed to clean up %v", err)
			}

			data := ReturnInfo{
				ImageName: imageName,
			}
			// info about built container
			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("Failed to marshal JSON: %v", err)
			}

			err = ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: d.CorrelationId,
					Body:          jsonData,
				})

			if err != nil {
				log.Fatalf("Failed to publish a message: %v", err)
			}
			d.Ack(true)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
