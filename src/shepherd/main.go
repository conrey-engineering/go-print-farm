package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"go.uber.org/zap"
)

func NewPrinter(db *gorm.DB, printer *pb.Printer) {
	printerDb := Printer{
		Name: printer.Name,
		APIConfig: PrinterAPIConfig{
			Type: "octoprint",
			Secret: printer.Api.Secret,
			Hostname: printer.Api.Hostname,
			Port: printer.Api.Port,
		},
	}

	db.Create(&printerDb)
}

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

func HandlePrinterTopicEvents(ctx context.Context, logger *zap.SugaredLogger, db *gorm.DB) {
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

		iterPrinter := eventMessage.Printer

		switch eventType := eventMessage.Type; eventType {
		case pb.PrinterEvent_CREATE:
			fmt.Println("Creating printer")
			NewPrinter(db, iterPrinter)
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

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()

	go func() {
		for { 
			_ = <-sigs
			done <- true
		}
	}()

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=password dbname=print_farm2"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Database migrations
	// db.AutoMigrate(&PrinterAPIType{})
	db.AutoMigrate(&PrinterAPIConfig{})
	db.AutoMigrate(&Printer{})

	go HandlePrinterTopicEvents(ctx, sugar, db)

	<-done

}