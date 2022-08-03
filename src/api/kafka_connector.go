package main

import (
	"github.com/segmentio/kafka-go"
)

type KafkaConnector struct {
	Brokers  []string
	MinBytes int
	MaxBytes int
}

func (w *KafkaConnector) newReaderConfig(topic string, partition int) kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers:   w.Brokers,
		Topic:     topic,
		Partition: partition,
		MinBytes:  w.MinBytes,
		MaxBytes:  w.MaxBytes,
	}
}

func (w *KafkaConnector) newReader(topic string, partition int) *kafka.Reader {
	config := w.newReaderConfig(topic, partition)
	return kafka.NewReader(config)
}

func (w *KafkaConnector) newWriterConfig(topic string, partition int) kafka.WriterConfig {
	return kafka.WriterConfig{
		Brokers: w.Brokers,
		Topic:   topic,
		// Partition: partition,
		// MinBytes:  w.MinBytes,
		// MaxBytes:  w.MaxBytes,
	}
}

func (w *KafkaConnector) newWriter(topic string, partition int) *kafka.Writer {
	config := w.newWriterConfig(topic, partition)
	return kafka.NewWriter(config)
}
