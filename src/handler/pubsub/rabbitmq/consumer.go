package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *rabbitmq) testConsumer1(q amqp.Queue) {
	msgs, err := r.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		r.log.Error().Err(err).Msg("Failed to register a consumer !!!")
	}

	// var forever chan struct{}

	go func() {
		for d := range msgs {
			r.log.Debug().Any("Received a message: ", string(d.Body)).Send()
		}
	}()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// <-forever
}
