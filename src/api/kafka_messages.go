package main

import (
	"sync"
)

type KafkaMessages struct {
	Messages []interface{}
	mu       sync.Mutex
}

func (m *KafkaMessages) Add(message interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Messages = append(m.Messages, message)
}
