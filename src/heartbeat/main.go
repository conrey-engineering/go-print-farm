package main

import (
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
	"github.com/segmentio/kafka-go"
	"context"
	// "time"
	"log"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"encoding/json"
)

func main() {

	topic := "printer_events"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "127.0.0.1:9092", topic, partition)

	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	examplePrinter := pb.Printer{
		Id: 5,
		Name: "Test Printer",
		Api: &pb.PrinterAPI{
			Type: pb.PrinterAPI_OCTOPRINT,
			Secret: "6879EBD309D34FA9B85FF8555A87B35E",
			Hostname: "localhost",
			Port: 80,
		},
	}

	// p := PrinterWrapper{
	// 	Printer: &examplePrinter,
	// }

	// fmt.Println(p)
	// fmt.Println(p.Poll())
	// fmt.Println(p.Status())

	// printerJson, err := json.Marshal(p)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	
	// Create example printer
	var event = pb.PrinterEvent{
		Type: pb.PrinterEvent_ONLINE,
		Printer: &examplePrinter,
	}
	data, err := json.Marshal(event)
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(string(data))
	_, err = conn.WriteMessages(
		kafka.Message{
			Value: data,
		},
		// kafka.Message{Value: printerJson},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		_ = <-sigs
		done <- true
	}()

	<-done

}