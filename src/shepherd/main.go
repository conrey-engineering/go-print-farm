package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
)

func generateKafkaReader(topic string) *kafka.Reader {
	logMsg := fmt.Sprintf("Created Kafka Reader for topic: %s", topic)
	fmt.Println(logMsg)
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"127.0.0.1:9092"},
		Topic: topic,
		Partition: 0,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func HandlePrinterTopicEvents(ctx context.Context) {
	rdr := generateKafkaReader("printer_events")
	rdr.SetOffset(0)
	fmt.Println("Listening for events...")
	for { 
		msg, err := rdr.ReadMessage(ctx)
		if err != nil {
			panic(err.Error())
		}

		var eventMessage = pb.PrinterEvent{}
		json.Unmarshal(msg.Value, &eventMessage)
		if eventMessage.Printer == nil {
			fmt.Println("Error: Printer not found in event message")
			continue
		}

		switch eventType := eventMessage.Type; eventType {
		case pb.PrinterEvent_CREATE:
			fmt.Println("Creating printer")
		case pb.PrinterEvent_DELETE:
			fmt.Println("Deleting printer")
		case pb.PrinterEvent_ERROR:
			fmt.Println("Marking printer as errored")
		case pb.PrinterEvent_OFFLINE:
			fmt.Println("Marking printer as offline")
		case pb.PrinterEvent_ONLINE:
			fmt.Println("Marking printer as online")
		default:
			fmt.Println("No idea")
		}
		printerJson, err := json.Marshal(eventMessage.Printer)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(string(printerJson))
	}
}

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	ctx := context.Background()

	go func() {
		for { 
			_ = <-sigs
			done <- true
		}
	}()

	go HandlePrinterTopicEvents(ctx)

	<-done

}