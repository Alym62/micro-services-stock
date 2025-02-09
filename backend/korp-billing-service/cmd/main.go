package main

import "github.com/Alym62/backend/korp-billing-service/internal/queue"

func main() {
	rabbitUrl := "amqp://guest:guest@localhost:5672/"
	queueName := "products.v1.invoice-event"
	exhangeName := "products.v1.product"
	consumer, err := queue.NewRabbitMQConsumer(rabbitUrl, exhangeName, queueName)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	go consumer.ConsumeMessages(rabbitUrl)

	select {}
}
