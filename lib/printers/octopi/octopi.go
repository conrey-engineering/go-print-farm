package octopi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Printer struct {
	Hostname string `json:"endpoint_url"`
	Port     int
	APIKey   string
}

type PrinterAPIVersion struct {
	Server string `json:"server"`
	Text   string `json:"text"`
	Api    string `json:"api"`
}

var (
	httpClient = http.Client{}
)

const (
	PRINTER_READY   = 0
	PRINTER_ERROR   = 1
	PRINTER_OFFLINE = 2
)

var PRINTER_STATUSES = map[int]string{
	0: "READY",
	1: "ERROR",
	2: "OFFLINE",
}

func doRequest(verb string, url string, apiKey string) (*http.Response, error) {
	req, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Api-Key", apiKey)

	return httpClient.Do(req)
}

func (p *Printer) Version() (string, error) {
	req_url := fmt.Sprintf("http://%s:%d/api/version", p.Hostname, p.Port)
	resp, err := doRequest("GET", req_url, p.APIKey)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var versionInfo = PrinterAPIVersion{}
	err = json.Unmarshal(body, &versionInfo)
	if err != nil {
		panic(err.Error())
	}

	return versionInfo.Server, nil

}

func (p *Printer) Status() (int, []byte, error) {
	req_url := fmt.Sprintf("http://%s:%d/api/printer?history=false&limit=1", p.Hostname, p.Port)
	resp, err := doRequest("GET", req_url, p.APIKey)
	if err != nil {
		return -1, nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, err
	}

	var statusMessage = map[string]string{}
	err = json.Unmarshal(body, &statusMessage)
	if err != nil {
		return -1, nil, err
	}

	for status, message := range statusMessage {
		switch s := status; s {
		case "error":
			return PRINTER_ERROR, []byte(message), nil
		}
	}
	return -1, nil, nil

}
