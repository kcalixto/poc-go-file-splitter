package main

import (
	"os"
	"time"
)

func main() {
	file := os.Getenv("FILE_PATH")
	if file == "" {
		panic("FILE_PATH env var is required")
	}

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
