package main

type Chunk struct {
	data []byte
	from int
	to   int
}

type ChunkHandler interface {
	Save(*Chunk) error
}
