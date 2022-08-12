package main

import (
	"context"
	"encoding/json"
	"github.com/conrey-engineering/go-print-farm/src/protobufs/print"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func printerHeartbeatMessageWatchdog(reader *kafka.Reader, messages KafkaMessages, logger *zap.SugaredLogger) {
	logger.Infow("Starting Printer heartbeat message watchdog")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			logger.Error(err.Error())
			break
		}
		messages.Add(string(msg.Value))
	}
}

func printerEventMessageWatchdog(reader *kafka.Reader, messages KafkaMessages, logger *zap.SugaredLogger) {
	logger.Infow("Starting Printer event message watchdog")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			logger.Error(err.Error())
		}

		var eventMessage = pb.PrinterEvent{}
		json.Unmarshal(msg.Value, &eventMessage)
		if eventMessage.Printer == nil {
			logger.Errorw("Error: Printer not found in event message")
			continue
		}
		messages.Add(string(msg.Value))
	}
}

func printEventMessageWatchdog(writer *kafka.Writer, c chan print.PrintRequestEvent, logger *zap.SugaredLogger) {
	logger.Infow("Starting Print event message watchdog")
	for {
		// For each request in the queue, publish a kafka message on the appropriate topic
		for eventMsg := range c {
			data, err := json.Marshal(eventMsg)
			if err != nil {
				logger.Errorw("Error marshaling JSON data from event message",
					"data", string(data),
					"error", err.Error(),
				)
			}

			err = writer.WriteMessages(context.Background(),
				kafka.Message{
					Topic: "print_events",
					Key:   []byte("print_request"),
					Value: data,
				},
			)

			if err != nil {
				logger.Errorw("Error writing kafka message",
					"message", string(data),
					"error", err.Error(),
				)
			}
		}
	}
}
