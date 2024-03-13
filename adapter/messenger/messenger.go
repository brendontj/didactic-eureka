package messenger

import (
	"context"
	"github.com/brendontj/didactic-eureka/infrastructure/rabbitmq"
)

type Adapter struct {
	*rabbitmq.Client
}

func NewAdapter(client *rabbitmq.Client) *Adapter {
	return &Adapter{client}
}

func (a *Adapter) DeclareQueue(name string) error {
	return a.QueueDeclare(name)
}

func (a *Adapter) Publish(ctx context.Context, queueName string, message []byte) error {
	return a.PublishMessage(ctx, queueName, message)
}
