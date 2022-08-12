package main

import (
	"context"
	"encoding/json"
	"fmt"
	libKafka "github.com/conrey-engineering/go-print-farm/lib/kafka"
	"github.com/conrey-engineering/go-print-farm/src/protobufs/print"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	KafkaReaderMinBytes = int(10e3) // 10KB
	KafkaReaderMaxBytes = int(10e6) // 10MB
	KafkaBrokers        = []string{"127.0.0.1:9092"}
	KafkaPartition      = 0

	kafkaConn = libKafka.KafkaConnector{
		Brokers: KafkaBrokers,
	}
)

func newPrinter(db *gorm.DB, printer *pb.Printer) {
	var (
		printerID, _     = uuid.Parse(printer.Id)
		printerApiConfig = PrinterAPIConfig{
			Type:     "octoprint",
			Secret:   printer.Api.Secret,
			Hostname: printer.Api.Hostname,
			Port:     printer.Api.Port,
		}
		printerDb = Printer{
			ID:        printerID,
			Name:      printer.Name,
			APIConfig: printerApiConfig,
		}
	)

	// If a printer of name printer.Name does not exist in the DB, create it
	go db.Debug().Create(&printerDb)
}

func handlePrinterTopicEvents(rdr *kafka.Reader, logger *zap.SugaredLogger, db *gorm.DB) {
	for {
		msg, err := rdr.ReadMessage(context.Background())
		if err != nil {
			panic(err.Error())
		}

		var eventMessage = pb.PrinterEvent{}
		json.Unmarshal(msg.Value, &eventMessage)
		if eventMessage.Printer == nil {
			logger.Errorw("Error: Printer not found in event message")
			continue
		}

		iterPrinter := eventMessage.Printer

		switch eventType := eventMessage.Type; eventType {
		case pb.PrinterEvent_CREATE:
			logger.Infow("Create printer",
				"printer_name", iterPrinter.Name,
			)
			newPrinter(db, iterPrinter)
		case pb.PrinterEvent_DELETE:
			logger.Infow("Deleting printer",
				"printer_name", iterPrinter.Name,
			)
		case pb.PrinterEvent_ERROR:
			logger.Infow("Marking printer as errored",
				"printer_name", iterPrinter.Name,
			)
		case pb.PrinterEvent_OFFLINE:
			logger.Infow("Marking printer as offline",
				"printer_name", iterPrinter.Name,
			)
		case pb.PrinterEvent_ONLINE:
			logger.Infow("Marking printer as online",
				"printer_name", iterPrinter.Name,
			)
		default:
			fmt.Println("No idea")
		}
	}
}

func handlePrintTopicEvents(rdr *kafka.Reader, logger *zap.SugaredLogger, db *gorm.DB) {
	for {
		msg, err := rdr.ReadMessage(context.Background())
		if err != nil {
			panic(err.Error())
		}

		var eventMessage = print.PrintRequestEvent{}
		json.Unmarshal(msg.Value, &eventMessage)
		if eventMessage.Request.Name == "" {
			logger.Errorw("Error: Print Request not found in event message")
			continue
		}

		var request = eventMessage.Request

		switch eventType := eventMessage.Type; eventType {
		case print.PrintRequestEvent_CREATE:
			logger.Infow("Create print request",
				"name", request.Name,
			)
			go db.Debug().Create(&PrintRequest{
				Name: request.Name,
			})
		default:
			fmt.Println("No idea")
		}
	}
}

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

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
	db.AutoMigrate(&PrinterAPIConfig{})
	db.AutoMigrate(&Printer{})
	db.AutoMigrate(&PrintRequest{})

	var (
		printEventReader   = kafkaConn.NewReader("print_events", 0)
		printerEventReader = kafkaConn.NewReader("printer_events", 0)
	)

	go handlePrinterTopicEvents(printerEventReader, sugar, db)
	go handlePrintTopicEvents(printEventReader, sugar, db)

	<-done

}
