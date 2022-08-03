package main

import (
	"fmt"
	pb "github.com/conrey-engineering/go-print-farm/src/protobufs/printer"
	"io"
	"net/http"
	"time"
)

type OctopiPrinter struct {
	*pb.Printer
}

func (p *OctopiPrinter) Poll() (*pb.PrinterStatus, error) {
	var status = pb.PrinterStatus{}
	req_url := fmt.Sprintf("http://%s:%d/api/printer?history=false&limit=1", p.Printer.Api.Hostname, p.Printer.Api.Port)
	req, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Api-Key", p.Printer.Api.Secret)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch statusCode := resp.StatusCode; statusCode {
	case 403:
		status.State = pb.PrinterStatus_ERROR
	case 409:
		status.State = pb.PrinterStatus_INACTIVE
		status.Message = string(body)
	case 200:
		status.State = pb.PrinterStatus_ACTIVE
	}

	p.Printer.Status = &status
	return p.Printer.Status, nil
}
