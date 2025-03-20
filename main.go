package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	file := os.Getenv("FILE_PATH")
	if file == "" {
		panic("FILE_PATH env var is required")
	}
	handleKind := os.Getenv("HANDLER_KIND")
	if handleKind == "" {
		panic("HANDLER_KIND env var is required")
	}
	var handler ChunkHandler
	switch handleKind {
	case "aws":
		handler = NewAWSS3ChunkHandler(context.Background(), time.Second*3)
	case "local":
		handler = LocalFileSystemHandler{}
	}
	if handler == nil {
		panic("Invalid handler kind")
	}

	maxWorkersStr := os.Getenv("MAX_WORKERS")
	if maxWorkersStr == "" {
		maxWorkersStr = "5"
	}
	maxWorkersInt, err := strconv.Atoi(maxWorkersStr)
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %v\n", r)
			fmt.Printf("panicked at %s\n", time.Now().String())
		}
	}()

	maxChunkSizeBytes := int(1024 * 1024 * 10)
	maxLines := 50000

	if err := NewFileBroker(
		FileBrokerConfig{
			nameType: NameTypeSequential,
			chunkBy:  ChunkByLines,
		},
		maxWorkersInt,
		time.Duration(time.Second*30),
		maxChunkSizeBytes,
		maxLines,
		handler,
	).Exec(file); err != nil {
		panic(err)
	}
}
