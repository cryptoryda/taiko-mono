package queue

import (
	"context"
	"errors"
	"sync"

	"github.com/taikoxyz/taiko-mono/packages/relayer/bindings/bridge"
)

var (
	ErrClosed = errors.New("queue connection closed")
)

type Queue interface {
	Start(ctx context.Context, queueName string) error
	Close(ctx context.Context)
	Publish(ctx context.Context, msg []byte) error
	Notify(ctx context.Context, wg *sync.WaitGroup) error
	Subscribe(ctx context.Context, msgs chan<- Message, wg *sync.WaitGroup) error
	Ack(ctx context.Context, msg Message) error
	Nack(ctx context.Context, msg Message) error
}

type QueueMessageSentBody struct {
	Event *bridge.BridgeMessageSent
	ID    int
}

type QueueMessageReceivedBody struct {
	Event *bridge.BridgeMessageReceived
	ID    int
}

type Message struct {
	Body     []byte
	Internal interface{}
}

type NewQueueOpts struct {
	Username      string
	Password      string
	Host          string
	Port          string
	PrefetchCount uint64
}
