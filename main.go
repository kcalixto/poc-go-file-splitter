package main

import "os"

func main() {
	file := os.Getenv("FILE_PATH")
	if file == "" {
		file = "big.csv"
	}

	if err := NewFileBroker(
		5,                    // maxCachedChunks
		int(1024*1024*10),    // 10MB - maxChunkSizeInBytes
		&AWSS3ChunkHandler{}, // chunkHandler
		// &LocalFileSystemHandler{}, // chunkHandler
	).Exec(file); err != nil {
		panic(err)
	}
}
