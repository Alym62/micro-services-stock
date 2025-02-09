package queue

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
	queue    string
}

func NewRabbitMQPublisher(url, exchange, queue string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		fmt.Printf("Erro ao conectar com o rabbit: %v", err)
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		fmt.Printf("Erro com o canal: %v", err)
		return nil, err
	}

	err = channel.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Erro ao criar a exchange: %v", err)
		return nil, err
	}

	_, err = channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Erro ao criar a fila: %v", err)
		return nil, err
	}

	err = channel.QueueBind(
		queue,
		queue,
		exchange,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Erro ao fazer o bind: %v", err)
		return nil, err
	}

	return &RabbitMQPublisher{conn: conn, channel: channel, exchange: exchange, queue: queue}, nil
}

func (p *RabbitMQPublisher) Publish(message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Erro com a message: %v", err)
		return err
	}

	err = p.channel.Publish(
		p.exchange,
		p.queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		fmt.Printf("Erro ao publicar a mensagem: %v", err)
		return err
	}

	fmt.Println("Mensagem publicada:", string(body))
	return nil
}

func (p *RabbitMQPublisher) Close() {
	p.channel.Close()
	p.conn.Close()
}
