package main

import (
	"context"
	"encoding/json"
	"fmt"
	libKafka "github.com/conrey-engineering/go-print-farm/lib/kafka"
	"github.com/conrey-engineering/go-print-farm/src/protobufs/print"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	KafkaReaderMinBytes = int(10e3) // 10KB
	KafkaReaderMaxBytes = int(10e6) // 10MB
	KafkaBrokers        = []string{"127.0.0.1:9092"}
	KafkaPartition      = 0

	// kafkaConn = newKafkaConnector(KafkaBrokers)
	kafkaConn = libKafka.KafkaConnector{
		Brokers: KafkaBrokers,
	}
	PrinterEventReader       = kafkaConn.NewReader("printer_events", 0)
	PrinterHeartbeatReader   = kafkaConn.NewReader("printer_heartbeats", 0)
	KafkaWriter              = kafkaConn.NewWriter(0)
	Logger, _                = zap.NewProduction()
	SugarLogger              = Logger.Sugar()
	printerEventMessages     = KafkaMessages{}
	printerHeartbeatMessages = KafkaMessages{}
	printFileQueue           = PrintRequestQueue{
		c: make(chan print.PrintRequestEvent),
	}
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func servePrinterHeartbeatLog(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(printerHeartbeatMessages.Messages)
	fmt.Fprintf(w, "%s", string(data))
}

func servePrinterEventLog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", printerEventMessages.Messages)
}

func servePrintCreate(w http.ResponseWriter, r *http.Request) {
	var (
		printRequest = print.PrintRequest{}
		body, _      = io.ReadAll(r.Body)
	)

	err := json.Unmarshal(body, &printRequest)
	if err != nil {
		SugarLogger.Errorw("Error converting data to PrintRequest",
			"error", err.Error(),
		)
	}

	SugarLogger.Infow("Print Create Request",
		"name", printRequest.Name,
	)

	printEvent := print.PrintRequestEvent{
		Action:  print.PrintRequestEvent_CREATE,
		Request: &printRequest,
	}

	go printFileQueue.Add(printEvent)
	w.WriteHeader(http.StatusCreated)
	return
}

func servePrinterCreate(w http.ResponseWriter, r *http.Request) {
	var (
		body, _      = io.ReadAll(r.Body)
		requestData  = map[string]string{}
		expectedKeys = []string{
			"printer_name",
			"hostname",
			"port",
			"api_token",
		}
	)

	json.Unmarshal(body, &requestData)

	for _, key := range expectedKeys {
		// Check if key from expectedKeys is in requestData
		if _, ok := requestData[key]; !ok {
			msg := fmt.Sprintf("missing key in request data: %s", key)

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(msg))
			return
		}
	}

	printerEventMessage := pb.PrinterEvent{
		Type: pb.PrinterEvent_CREATE,
		Printer: &pb.Printer{
			Name: requestData["printer_name"],
			Api: &pb.PrinterAPI{
				Type:     pb.PrinterAPI_OCTOPRINT,
				Secret:   requestData["api_token"],
				Hostname: requestData["hostname"],
				Port:     80,
			},
		},
	}

	msgData, _ := json.Marshal(printerEventMessage)
	err := KafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Topic: "printer_events",
			Value: msgData,
		},
	)
	if err != nil {
		SugarLogger.Errorw("Problem sending kafka message",
			"message", string(msgData),
		)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SugarLogger.Infow("Handling request",
			"method", r.Method,
			"remote_addr", r.RemoteAddr,
			"request_url", r.URL,
		)
		next.ServeHTTP(w, r)
	})
}

func main() {
	go printerEventMessageWatchdog(PrinterEventReader, printerEventMessages, SugarLogger)
	go printerHeartbeatMessageWatchdog(PrinterHeartbeatReader, printerHeartbeatMessages, SugarLogger)
	go printEventMessageWatchdog(KafkaWriter, printFileQueue.c, SugarLogger)

	var wait time.Duration
	defer Logger.Sync()

	mux := mux.NewRouter()
	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}
	mux.Use(loggingMiddleware)
	mux.HandleFunc("/printers/create", servePrinterCreate).Methods("POST")
	mux.HandleFunc("/printers/events", servePrinterEventLog)
	mux.HandleFunc("/printers/heartbeats", servePrinterHeartbeatLog)
	mux.HandleFunc("/print/create", servePrintCreate).Methods("POST")
	// mux.HandleFunc("/print/{id}/upload", serverPrintFileUpload).Methods("POST")
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

	// // mux.Handle("/", hello)
	// mux.Handle("/printers/events",
	//     loggingMiddleware(sugar,
	//         http.HandlerFunc(servePrinterEventLog),
	//     ),
	// )
	// http.HandleFunc("/hello", hello)
	// http.HandleFunc("/headers", headers)
	// http.HandleFunc("/printers/add", servePrinterAdd)

	// http.ListenAndServe(":8090", mux)
}
