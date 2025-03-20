package main

import (
	"os"
	"testing"
	"time"
)

type MockChunkHandler struct{}

func (h *MockChunkHandler) Save(c *Chunk) error {
	time.Sleep(3000 * time.Millisecond)
	return nil
}

func BenchmarkFileBroker(b *testing.B) {
	cfg := FileBrokerConfig{
		nameType: NameTypeSequential,
	}
	mockChunkHandler := &MockChunkHandler{}

	broker := NewFileBroker(cfg, 5, time.Second, int(1024*1024*10), 5, mockChunkHandler)
	start := time.Now()
	b.ResetTimer()
	b.ReportAllocs()
	file := os.Getenv("FILE_PATH")
	if file == "" {
		file = "big.csv"
	}
	broker.Exec(file)
	since := time.Since(start)
	if since.Minutes() > 15 {
		b.Fail()
	}
	b.Log("execTime: ", since.String())
}
