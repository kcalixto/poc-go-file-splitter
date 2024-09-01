package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type FileBroker struct {
	wg                  *sync.WaitGroup
	chunksChan          chan *Chunk
	errsChan            chan error
	maxChunkSizeInBytes int
	chunkHandler        ChunkHandler
}

func NewFileBroker(
	maxCachedChunks int,
	maxChunkSizeInBytes int,
	chunkHandler ChunkHandler,
) *FileBroker {
	return &FileBroker{
		wg:                  &sync.WaitGroup{},
		chunksChan:          make(chan *Chunk, maxCachedChunks),
		errsChan:            make(chan error, maxCachedChunks),
		maxChunkSizeInBytes: maxChunkSizeInBytes,
		chunkHandler:        chunkHandler,
	}
}

func (b *FileBroker) Exec(fileName string) error {
	reader, err := b.openFile(fileName)
	if err != nil {
		return err
	}

	go b.save()

	h, err := reader.ReadString('\n') // read header
	if err != nil {
		return err
	}
	header := []byte(h)

	avgLineSizeBytes := 1024 * 1 // 1KB
	lineCounter := 1
	lastSavedLine := 0
	buf := bytes.NewBuffer(make([]byte, 0, b.maxChunkSizeInBytes+(avgLineSizeBytes)))
	buf.Write(header)

	fmt.Printf("Started execution at %s\n", time.Now().Format(time.TimeOnly))
	for {
		if err := b.read(reader, buf); err != nil {
			if err == io.EOF {
				b.chunksChan <- &Chunk{
					data: buf.Bytes(),
					from: lastSavedLine + 1,
					to:   lineCounter,
				}
				// fmt.Printf("sending file from %d to %d\n", lastSavedLine+1, lineCounter)
				buf = new(bytes.Buffer)
				lastSavedLine = lineCounter

				break
			}
			return err
		}

		if buf.Len() >= b.maxChunkSizeInBytes {
			b.chunksChan <- &Chunk{
				data: buf.Bytes(),
				from: lastSavedLine + 1,
				to:   lineCounter,
			}
			// fmt.Printf("sending file from %d to %d\n", lastSavedLine+1, lineCounter)
			buf = bytes.NewBuffer(make([]byte, 0, b.maxChunkSizeInBytes+avgLineSizeBytes))
			buf.Write(header)
			lastSavedLine = lineCounter
		}

		lineCounter++
	}
	close(b.chunksChan)
	b.wg.Wait()
	close(b.errsChan)

	for err := range b.errsChan {
		fmt.Println("Error saving file: ", err)
	}

	fmt.Println("All files saved")
	return nil
}

func (b *FileBroker) openFile(name string) (*bufio.Reader, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(file), nil
}

func (b *FileBroker) read(in *bufio.Reader, out *bytes.Buffer) error {
	str, err := in.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			if str != "" {
				out.Write([]byte(str))
			}
			return io.EOF
		}
		return err
	}

	out.Write([]byte(str))
	return nil
}

func (b *FileBroker) save() {
	for f := range b.chunksChan {
		b.wg.Add(1)
		go func(f *Chunk) {
			defer b.wg.Done()
			err := b.chunkHandler.Save(f)
			if err != nil {
				b.errsChan <- err
			}
		}(f)
	}
}
