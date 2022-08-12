package main

import (
	"context"
	"encoding/json"
	heartbeat "github.com/conrey-engineering/go-print-farm/src/protobufs/heartbeat"
	"github.com/segmentio/kafka-go"
)

func processHeartbeats(writer *kafka.Writer, heartbeatChan <-chan heartbeat.PrinterHeartbeat) {
	for {
		for heartbeat := range heartbeatChan {
			heartbeatJson, _ := json.Marshal(heartbeat)
			msg := kafka.Message{
				Topic: "printer_heartbeats",
				Key:   []byte("heartbeat"),
				Value: heartbeatJson,
			}
			err := writer.WriteMessages(context.Background(), msg)
			if err != nil {
				sugarLogger.Errorw("Error writing heartbeat message to kafka",
					"error", err.Error(),
				)
			} else {
				sugarLogger.Infow("Published heartbeat message to kafka",
					"info", string(msg.Value),
				)
			}
		}
	}
}
