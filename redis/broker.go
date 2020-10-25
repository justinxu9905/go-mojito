package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Broker struct {
	client *redis.Client
}

func New() Broker {
	return Broker{
		redis.NewClient(&redis.Options{
			Addr:		"localhost:6379",
			Password:	"",
			DB:			0,
		}),
	}
}

var RB = New()

func (b Broker) SendTask(key string, value string) (err error) {
	err = b.client.LPush(ctx, key, value).Err()
	return err
}

func (b Broker) ReceiveTask(key string) string {
	result, _ := b.client.RPop(ctx, key).Result()
	return result
}