package gateway

import "context"

type MessageGateway interface {
	Publish(ctx context.Context, queueName string, message []byte) error
}
