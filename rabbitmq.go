package gowaitfordeps

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type RabbitMQConfig struct {
	User     string
	Password string
	Host     string
	Port     string
}

// WaitForRabbitMQ waits for RabbitMQ to be up and running
func WaitForRabbitMQ(config RabbitMQConfig) {

	rabbitmqDSN := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.User,
		config.Password,
		config.Host,
		config.Port,
	)

	for {
		waitRabbitMQ, err := amqp.Dial(rabbitmqDSN)
		if err == nil {
			e := waitRabbitMQ.Close()
			if e != nil {
				return
			}
			break
		}

		log.Println("RabbitMQ connection error, will retry:", err)
		time.Sleep(5 * time.Second)
	}

	log.Println("RabbitMQ connection successful!")

}
