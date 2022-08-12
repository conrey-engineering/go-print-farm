package main

import (
	"github.com/conrey-engineering/go-print-farm/src/protobufs/print"
)

type PrintRequestQueue struct {
	c chan print.PrintRequestEvent
}

func (q *PrintRequestQueue) Add(p print.PrintRequestEvent) {
	q.c <- p
}
