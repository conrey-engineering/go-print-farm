package main

import (
	// "encoding/json"
	"fmt"
	"github.com/conrey-engineering/go-print-farm/lib/printers/octopi"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
)

type PrinterWrapper struct {
	*pb.Printer
}

// func (p *PrinterWrapper) Status() string {
// 	out, err := json.Marshal(p.Printer.Status)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return string(out)
// }

func octoprintPoll(p *PrinterWrapper) (pb.PrinterStatus, error) {
	var printerStatus = pb.PrinterStatus{}

	var printer = octopi.Printer{
		Hostname: p.Printer.Api.Hostname,
		Port:     int(p.Printer.Api.Port),
		APIKey:   p.Printer.Api.Secret,
	}
	status, message, err := printer.Status()
	if err != nil {
		return printerStatus, err
	}

	printerStatus.Message = string(message)

	switch status {
	case octopi.PRINTER_ERROR:
		printerStatus.State = pb.PrinterStatus_ERROR
	case octopi.PRINTER_OFFLINE:
		printerStatus.State = pb.PrinterStatus_INACTIVE
	case octopi.PRINTER_READY:
		printerStatus.State = pb.PrinterStatus_ACTIVE
	}

	p.Printer.Status = &printerStatus

	return *p.Printer.Status, nil
}

func (p *PrinterWrapper) Poll() (pb.PrinterStatus, error) {
	var printerState = pb.PrinterStatus{}

	switch api := p.Printer.Api.Type; api {
	case pb.PrinterAPI_OCTOPRINT:
		return octoprintPoll(p)
	default:
		fmt.Println("No API")
	}

	return printerState, nil

}

func (p *PrinterWrapper) Version() string {
	switch api := p.Printer.Api.Type; api {
	case pb.PrinterAPI_OCTOPRINT:
		var printer = octopi.Printer{
			Hostname: p.Printer.Api.Hostname,
			Port:     int(p.Printer.Api.Port),
			APIKey:   p.Printer.Api.Secret,
		}
		vers, err := printer.Version()
		if err != nil {
			panic(err.Error())
		}
		return vers
	default:
		fmt.Println("No API")
	}
	return ""
}
