package main

import (
	"context"
	"encoding/json"
	libKafka "github.com/conrey-engineering/go-print-farm/lib/kafka"
	heartbeat "github.com/conrey-engineering/go-print-farm/src/protobufs/heartbeat"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	logger, _           = zap.NewProduction()
	sugarLogger         = logger.Sugar()
	heartbeats          = make(chan heartbeat.PrinterHeartbeat)
	printers            []PrinterWrapper
	KafkaReaderMinBytes = int(10e3) // 10KB
	KafkaReaderMaxBytes = int(10e6) // 10MB
	KafkaBrokers        = []string{"127.0.0.1:9092"}
	KafkaPartition      = 0

	kafkaConn = libKafka.KafkaConnector{
		Brokers: KafkaBrokers,
	}
)

func pollPrinter(printer PrinterWrapper) {
	var heartbeat_result heartbeat.PrinterHeartbeat_Result

	printer.Poll()

	switch hb_type := printer.Printer.Status.State; hb_type {
	case pb.PrinterStatus_ACTIVE:
		heartbeat_result = heartbeat.PrinterHeartbeat_SUCCESS
	default:
		heartbeat_result = heartbeat.PrinterHeartbeat_FAILURE
	}

	heartbeats <- heartbeat.PrinterHeartbeat{
		Result:    heartbeat_result,
		PrinterId: printer.Printer.Id,
		Message:   printer.Printer.Status.Message,
	}

}

func pollPrinters() {
	for _, printer := range printers {
		go pollPrinter(printer)
	}
}

func processHeartbeats(writer *kafka.Writer, heartbeatChan <-chan heartbeat.PrinterHeartbeat) {
	for {
		for heartbeat := range heartbeatChan {
			heartbeatJson, _ := json.Marshal(heartbeat)
			msg := kafka.Message{
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

func main() {
	defer logger.Sync()
	examplePrinter := PrinterWrapper{&pb.Printer{
		Id:   "a40959ab-8b96-46ea-b8c7-f0cf169ff602",
		Name: "Test Printer",
		Api: &pb.PrinterAPI{
			Type:     pb.PrinterAPI_OCTOPRINT,
			Secret:   "6879EBD309D34FA9B85FF8555A87B35E",
			Hostname: "localhost",
			Port:     80,
		},
	}}

	printers = append(printers, examplePrinter)

	// goroutine for polling all printers and sleeping for 10s
	go func() {
		for {
			pollPrinters()
			time.Sleep(time.Second * 10)
		}
	}()

	sugarLogger.Infow("Creating kafka writer",
		"topic", "printer_heartbeats",
		"partition", 0,
	)
	writer := kafkaConn.NewWriter("printer_heartbeats", 0)

	// Watch the `heartbeats` chan for heartbeats, publish to Kafka as they occur]
	go processHeartbeats(writer, heartbeats)

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		_ = <-sigs
		done <- true
	}()

	<-done

}
