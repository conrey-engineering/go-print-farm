package main

import (
	heartbeat "github.com/conrey-engineering/go-print-farm/src/protobufs/heartbeat"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
)

func pollPrinter(printer PrinterWrapper) {
	var heartbeat_result heartbeat.PrinterHeartbeat_Result

	_, err := printer.Poll()
	if err != nil {
		panic(err.Error())
	}

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
