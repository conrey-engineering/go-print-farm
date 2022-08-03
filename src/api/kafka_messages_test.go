package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var msgs KafkaMessages
	msgs.Add("test")
}

func BenchmarkAdd(b *testing.B) {
	var msgs KafkaMessages
	for i := 0; i < b.N; i++ {
		msgs.Add("test")
	}
}
