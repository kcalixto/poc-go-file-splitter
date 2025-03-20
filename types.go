package main

type Chunk struct {
	name string
	data []byte
	from int
	to   int
}

type ChunkHandler interface {
	Save(*Chunk) error
}
