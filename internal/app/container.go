package app

import (
	"food-delivery-backend/internal/grpc/client"
	"food-delivery-backend/pkg/config"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Container struct {
	Config      *config.Config
	Logger      *zap.Logger
	DB          *sqlx.DB
	Redis       *redis.Client
	KafkaWriter *kafka.Writer
	OrderClient *client.OrderServiceClient
}
