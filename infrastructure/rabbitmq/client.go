package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

type Config struct {
	User string
	Pass string
	Host string
	Port string
}

func NewClient(c Config) (*Client, error) {
	client := &Client{}
	if err := client.connect(c); err != nil {
		return client, err
	}

	return client, nil
}

func (c *Client) QueueDeclare(name string) error {
	_, err := c.Channel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (c *Client) PublishMessage(ctx context.Context, queueName string, message []byte) error {
	return c.Channel.PublishWithContext(
		ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
}

func (c *Client) Close() {
	defer func() {
		_ = c.Conn.Close()
		_ = c.Channel.Close()
	}()
}

func (c *Client) connect(config Config) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.User,
		config.Pass,
		config.Host,
		config.Port))
	if err != nil {
		return err
	}

	c.Conn = conn

	ch, err := c.openChannel()
	if err != nil {
		return err
	}
	c.Channel = ch
	return nil
}

func (c *Client) openChannel() (*amqp.Channel, error) {
	ch, err := c.Conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}
