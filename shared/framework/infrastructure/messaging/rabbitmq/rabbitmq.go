package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConn struct {
	*amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQConn(uri, queue string) (*RabbitMQConn, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	_, err = ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return &RabbitMQConn{conn, ch}, nil
}

func (r *RabbitMQConn) Close() error {
	if err := r.Channel.Close(); err != nil {
		return err
	}
	return r.Connection.Close()
}
