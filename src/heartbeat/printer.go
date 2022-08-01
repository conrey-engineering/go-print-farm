package main

import (
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
	"encoding/json"
	"fmt"
)

type PrinterWrapper struct {
	*pb.Printer
}

func (p *PrinterWrapper) Status() string {
	out, err := json.Marshal(p.Printer.Status)
	if err != nil {
		panic(err.Error())
	}
	return string(out)
}

func (p *PrinterWrapper) Poll() bool {
	switch api := p.Printer.Api.Type; api {
	case pb.PrinterAPI_OCTOPRINT:
		var printer = OctopiPrinter{p.Printer}
		_, err := printer.Poll()
		if err != nil {
			panic(err.Error())
		}
	default:
		fmt.Println("No API")
	}

	return true

}