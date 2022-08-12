package printer_helper

import (
	"encoding/json"
	"fmt"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
)

type PrinterHelper struct {
	*pb.Printer
}

func (p *PrinterHelper) Status() string {
	out, err := json.Marshal(p.Printer.Status)
	if err != nil {
		panic(err.Error())
	}
	return string(out)
}

func (p *PrinterHelper) Poll() bool {
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
