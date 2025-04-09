package pubsub

import "context"

type SubMessage struct {
	From string
	Data string
}

type PubSubRepository interface {
	Publish(ctx context.Context, key string, payload string) error
	Subscribe(ctx context.Context, keys ...string) (<-chan *SubMessage, error)
}
