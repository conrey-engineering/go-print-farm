package main

import (
	"fmt"
	// "github.com/segmentio/kafka-go"
	//  "github.com/conrey-engineering/go-print-farm/lib/kafka"
	"context"
	"encoding/json"
	"github.com/conrey-engineering/go-print-farm/lib/kafka"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
	"github.com/google/uuid"
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

	kafkaConn = kafka.KafkaConnector{
		Brokers: KafkaBrokers,
	}
	// PrinterEventReader = kafkaConn.newReader("printer_events", 0)
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
	db.Debug().FirstOrCreate(&printerDb)
}

func handlePrinterTopicEvents(ctx context.Context, logger *zap.SugaredLogger, db *gorm.DB, kafka_topic string) {
	topic := "printer_events"
	partition := 0

	logger.Debugw("Creating kafka reader",
		"topic", topic,
		"partition", partition,
	)
	rdr := kafkaConn.NewReader(topic, partition)
	logger.Debugw("Created kafka reader",
		"topic", topic,
		"partition", partition,
	)

	rdr.SetOffset(0)
	logger.Infow("Listening for messages",
		"kafka_topic", kafka_topic,
	)

	for {
		msg, err := rdr.ReadMessage(ctx)
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

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	ctx := context.Background()

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
	// db.AutoMigrate(&PrinterAPIType{})
	db.AutoMigrate(&PrinterAPIConfig{})
	db.AutoMigrate(&Printer{})

	go handlePrinterTopicEvents(ctx, sugar, db, "printer_events")

	<-done

}
