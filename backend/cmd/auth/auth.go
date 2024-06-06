package main

import (
	"backend/db"
	"backend/internal/auth"
	"backend/internal/rabbitmq"
	"github.com/streadway/amqp"
	"log"
)

type AuthService struct {
	handler auth.Handler
}

func newAuthService() *AuthService {
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	authRepo := auth.NewRepository(conn.GetClient())
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	return &AuthService{authHandler}
}

func processMessage(service *AuthService, msg amqp.Delivery) string {
	// C
	return ""
}

func main() {
	service := newAuthService()
	service.handler.Login(auth.LoginReq{
		Username: "test",
		Password: "test",
	})
	conn, ch, err := rabbitmq.SetupRabbitMQ()
	defer func(conn *amqp.Connection) {
		err = conn.Close()
		if err != nil {
			log.Fatal("Error closing connection: ", err.Error())
		}
	}(conn)
	defer func(ch *amqp.Channel) {
		err = ch.Close()
		if err != nil {
			log.Fatal("Error closing channel: ", err.Error())
		}
	}(ch)

	msgs, err := rabbitmq.ConsumeMessages(ch, "auth_queue")
	if err != nil {
		log.Fatal("Error consuming messages: ", err.Error())
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			//log.Printf("Received a message: %s", d.Body)
			result := processMessage(service, msg)

			err = rabbitmq.PublishMessages(ch, "finished_queue", result)
			if err != nil {
				log.Println("Error publishing message: ", err.Error())
				continue
			}
			log.Println("Message processed: ", result)
		}
	}()

	<-forever
}
