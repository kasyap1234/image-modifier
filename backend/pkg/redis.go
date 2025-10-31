package pkg

import (
	"github.com/hibiken/asynq"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type AsynqRedisClient struct {
	client *asynq.Client
}

func NewAsynqRedisClient(redisConfig *RedisConfig) (*AsynqRedisClient, error) {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	return &AsynqRedisClient{
		client: client,
	}, nil
}

func (ac *AsynqRedisClient) GetClient() *asynq.Client {
	return ac.client
}

func (ac *AsynqRedisClient) Close() error {
	return ac.client.Close()
}