package db

import (
	"sync"
	"time"
)

var (
	m *sync.RWMutex
)

func ResetMetric() {
	for {
		READ = WRITE
		WRITE = 0
		time.Sleep(60 * time.Second	)
	}
}