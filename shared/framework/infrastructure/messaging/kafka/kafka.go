package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaConnection struct {
	Writer *kafka.Writer
	Reader *kafka.Reader
}

func NewKafkaConnection(brokers []string, topic, groupID string) (*KafkaConnection, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	return &KafkaConnection{
		Writer: writer,
		Reader: reader,
	}, nil
}

func (kc *KafkaConnection) Close() error {
	if err := kc.Writer.Close(); err != nil {
		return err
	}
	return kc.Reader.Close()
}

func (kc *KafkaConnection) ReadMessage(ctx context.Context) (kafka.Message, error) {
	return kc.Reader.ReadMessage(ctx)
}

func (kc *KafkaConnection) WriteMessage(ctx context.Context, msg kafka.Message) error {
	return kc.Writer.WriteMessages(ctx, msg)
}

func (kc *KafkaConnection) CommitMessages(ctx context.Context, msgs ...kafka.Message) error {
	return kc.Reader.CommitMessages(ctx, msgs...)
}
