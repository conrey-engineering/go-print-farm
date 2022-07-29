package main

import (
    "fmt"
	pb "src/protobufs/printer"
)

func main() {

	buf := new(bytes.Buffer)

	p := pb.Printer{
		Id: 1,
		Name: "Test Printer",
		PrinterStatus: 0,
	}
	fmt.Println("test")
}