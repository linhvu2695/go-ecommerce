package global

import (
	"database/sql"
	"go-ecommerce/pkg/logger"
	"go-ecommerce/pkg/settings"
	"net/smtp"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	Config        settings.Config
	Logger        *logger.LoggerZap
	Db            *sql.DB
	Redis         *redis.Client
	SmtpAuth      *smtp.Auth
	KafkaProducer *kafka.Writer
)
