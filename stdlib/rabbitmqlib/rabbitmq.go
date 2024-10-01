package rabbitmqlib

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

type Options struct {
	Host     string
	Port     string
	User     string
	Password preference.EncryptedValue
}

func Init(log zerolog.Logger, opt Options) *amqp.Channel {
	strConn := fmt.Sprintf("amqp://%s:%s@%s:%s/", opt.User, opt.Password, opt.Host, opt.Port)
	conn, err := amqp.Dial(strConn)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to connect to RabbitMQ !!!")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to open a channel !!!")
	}

	log.Debug().Msg(fmt.Sprintf("RabbitMQ status: OK"))

	return ch
}
