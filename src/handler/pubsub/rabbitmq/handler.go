package rabbitmq

import (
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

var once = &sync.Once{}

type rabbitmq struct {
	log     zerolog.Logger
	channel *amqp.Channel
}

func Init(log zerolog.Logger, c *amqp.Channel) {
	once.Do(func() {
		rmq := &rabbitmq{
			log:     log,
			channel: c,
		}

		rmq.Serve()
	})
}

func (r *rabbitmq) Serve() {
	q, err := r.channel.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		r.log.Error().Err(err).Msg("Failed to declare a queue !!!")
	}

	r.testConsumer1(q)
}
