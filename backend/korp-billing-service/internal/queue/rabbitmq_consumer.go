package queue

import (
	"fmt"
	"time"

	"github.com/Alym62/backend/korp-billing-service/pkg"
	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
	queue    string
}

func NewRabbitMQConsumer(url, exchange, queue string) (*RabbitMQConsumer, error) {
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

	return &RabbitMQConsumer{conn: conn, channel: channel, exchange: exchange, queue: queue}, nil
}

func (c *RabbitMQConsumer) ConsumeMessages(url string) {
	for {
		if c.channel == nil {
			fmt.Println("Canal fechado. Tentando reconectar...")
			err := c.Reconnect(url)
			if err != nil {
				fmt.Println("Falha ao reconectar. Tentando novamente em 5s...")
				time.Sleep(5 * time.Second)
				continue
			}
		}

		messages, err := c.channel.Consume(
			c.queue,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			fmt.Printf("Erro ao consumir mensagens: %v\n", err)
			c.channel = nil
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("Consumindo mensagens da fila:", c.queue)

		for message := range messages {
			fmt.Println("Mensagem recebida:", string(message.Body))
			err := pkg.GeneratePDF(message.Body)
			if err != nil {
				fmt.Printf("Erro ao processar NF, enviando para DLQ: %v\n", err)
			}
		}
	}
}

func (c *RabbitMQConsumer) Close() {
	c.channel.Close()
	c.conn.Close()
}

func (c *RabbitMQConsumer) Reconnect(url string) error {
	fmt.Println("Tentando reconectar ao RabbitMQ...")

	if c.conn != nil {
		c.conn.Close()
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		fmt.Printf("Erro ao reconectar ao RabbitMQ: %v\n", err)
		return err
	}
	channel, err := conn.Channel()
	if err != nil {
		fmt.Printf("Erro ao abrir canal na reconexão: %v\n", err)
		return err
	}

	c.conn = conn
	c.channel = channel
	fmt.Println("Reconexão bem-sucedida!")
	return nil
}
