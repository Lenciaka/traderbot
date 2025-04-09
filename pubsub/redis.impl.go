package pubsub

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type RedisPubSub struct {
	client *redis.Client
}

func NewRedisPudSub(client *redis.Client) *RedisPubSub {
	return &RedisPubSub{
		client: client,
	}
}

func (pubsub *RedisPubSub) Publish(ctx context.Context, key string, payload string) error {
	_, err := pubsub.client.Publish(ctx, key, payload).Result()
	log.Info().Msgf("publish payload to %s", key)
	return err
}

func (pubsub *RedisPubSub) Subscribe(ctx context.Context, keys ...string) (<-chan *SubMessage, error) {
	defer log.Info().Msgf("Subscribe %v was down", keys)
	msg := make(chan *SubMessage, 10)
	ch := pubsub.client.Subscribe(ctx, keys...).Channel()
	go func() {
		select {
		case <-ctx.Done():
			break
		case message := <-ch:
			msg <- &SubMessage{
				From: message.Channel,
				Data: message.Payload,
			}
		}
	}()
	return msg, nil
}
