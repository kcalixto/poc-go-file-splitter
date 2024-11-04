package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file := os.Getenv("FILE_PATH")
	if file == "" {
		panic("FILE_PATH env var is required")
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %v\n", r)
			fmt.Printf("panicked at %s\n", time.Now().String())
		}
	}()
	if err := NewFileBroker(
		5,                             // maxConcurrentWorkers
		time.Duration(time.Second*30), // delayUntilInvokeNewWorker
		int(1024*1024*10),             // 10MB - maxChunkSizeInBytes
		&AWSS3ChunkHandler{},          // chunkHandler
		// &LocalFileSystemHandler{}, // chunkHandler
	).Exec(file); err != nil {
		panic(err)
	}
}
