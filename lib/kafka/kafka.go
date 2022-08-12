package kafka

import (
	"github.com/segmentio/kafka-go"
)

type KafkaConnector struct {
	Brokers  []string
	MinBytes int
	MaxBytes int
}

func (w *KafkaConnector) newReaderConfig(topic string, partition int) kafka.ReaderConfig {
	// var logger KafkaLogger
	return kafka.ReaderConfig{
		Brokers:   w.Brokers,
		Topic:     topic,
		Partition: partition,
		MinBytes:  w.MinBytes,
		MaxBytes:  w.MaxBytes,
		// Logger:    kafka.LoggerFunc(logger.Log),
	}
}

func (w *KafkaConnector) NewReader(topic string, partition int) *kafka.Reader {
	config := w.newReaderConfig(topic, partition)
	return kafka.NewReader(config)
}

// func (w *KafkaConnector) newWriterConfigWithTopic(topic string, partition int) kafka.WriterConfig {
// 	return kafka.WriterConfig{
// 		Brokers: w.Brokers,
// 		Topic:   topic,
// 		// Partition: partition,
// 		// MinBytes:  w.MinBytes,
// 		// MaxBytes:  w.MaxBytes,
// 	}
// }

// func (w *KafkaConnector) NewWriterWithTopic(topic string, partition int) *kafka.Writer {
// 	config := w.newWriterConfigWithTopic(topic, partition)
// 	return kafka.NewWriter(config)
// }

func (w *KafkaConnector) NewWriter(partition int) *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: w.Brokers,
	})

	writer.AllowAutoTopicCreation = true

	return writer
}
