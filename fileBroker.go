package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

type NameTypeEnum int

const (
	NameTypeUUID NameTypeEnum = iota + 1
	NameTypeSequential
)

type ChunkByEnum int

const (
	ChunkBySize ChunkByEnum = iota + 1
	ChunkByLines
)

type FileBrokerConfig struct {
	nameType NameTypeEnum
	chunkBy  ChunkByEnum
}

func (c FileBrokerConfig) Name(chunkCounter ...int) string {
	switch c.nameType {
	case NameTypeUUID:
		return uuid.New().String()
	case NameTypeSequential:
		if len(chunkCounter) > 0 {
			return fmt.Sprintf("%05d", chunkCounter[0])
		}
		return time.Now().Format("2006-01-02T15:04:05")
	default:
		panic("invalid name type")
	}
}

type FileBroker struct {
	cfg                       FileBrokerConfig
	wg                        *sync.WaitGroup
	maxConcurrentWorkers      int
	delayUntilInvokeNewWorker time.Duration
	chunksChan                chan *Chunk
	errsChan                  chan error
	maxChunkSizeInBytes       int
	maxLines                  int
	chunkHandler              ChunkHandler
}

func NewFileBroker(
	cfg FileBrokerConfig,
	maxConcurrentWorkers int,
	delayUntilInvokeNewWorker time.Duration,
	maxChunkSizeInBytes int,
	maxLines int,
	chunkHandler ChunkHandler,
) *FileBroker {
	return &FileBroker{
		cfg:                  cfg,
		wg:                   &sync.WaitGroup{},
		maxConcurrentWorkers: maxConcurrentWorkers,
		chunksChan:           make(chan *Chunk),
		errsChan:             make(chan error),
		maxChunkSizeInBytes:  maxChunkSizeInBytes,
		maxLines:             maxLines,
		chunkHandler:         chunkHandler,
	}
}

func (b *FileBroker) Exec(fileName string) error {
	reader, err := b.openFile(fileName)
	if err != nil {
		return err
	}

	h, err := reader.ReadString('\n') // read header
	if err != nil {
		return err
	}
	header := []byte(h)

	avgLineSizeBytes := 1024 * 1 // 1KB
	lineCounter := 1
	lastSavedLine := 0
	chunkCounter := 0
	buf := bytes.NewBuffer(make([]byte, 0, b.maxChunkSizeInBytes+(avgLineSizeBytes)))
	buf.Write(header)

	fmt.Printf("Started execution at %s\n", time.Now().Format(time.TimeOnly))
	endChan := make(chan struct{})
	go func() {
		for {
			if err := b.read(reader, buf); err != nil {
				if err == io.EOF {
					b.chunksChan <- &Chunk{
						name: b.cfg.Name(chunkCounter),
						data: buf.Bytes(),
						from: lastSavedLine + 1,
						to:   lineCounter,
					}
					// fmt.Printf("sending file from %d to %d\n", lastSavedLine+1, lineCounter)
					buf = new(bytes.Buffer)
					lastSavedLine = lineCounter

					break
				}
				panic(err) // file reading error
			}

			switch b.cfg.chunkBy {
			case ChunkBySize:
				if buf.Len() >= b.maxChunkSizeInBytes {
					b.chunksChan <- &Chunk{
						name: b.cfg.Name(chunkCounter),
						data: buf.Bytes(),
						from: lastSavedLine + 1,
						to:   lineCounter,
					}
					chunkCounter++
					buf = bytes.NewBuffer(make([]byte, 0, b.maxChunkSizeInBytes+avgLineSizeBytes))
					buf.Write(header)
					lastSavedLine = lineCounter
				}
			case ChunkByLines:
				if lineCounter%b.maxLines == 0 {
					b.chunksChan <- &Chunk{
						name: b.cfg.Name(chunkCounter),
						data: buf.Bytes(),
						from: lastSavedLine + 1,
						to:   lineCounter,
					}
					chunkCounter++
					buf = bytes.NewBuffer(make([]byte, 0, b.maxChunkSizeInBytes+avgLineSizeBytes))
					buf.Write(header)
					lastSavedLine = lineCounter
				}
				// fmt.Printf("sending file from %d to %d\n", lastSavedLine+1, lineCounter)
			}

			lineCounter++
		}
		close(b.chunksChan)
		b.wg.Wait()
		close(b.errsChan)

		for err := range b.errsChan {
			fmt.Println("Error saving file: ", err)
		}

		endChan <- struct{}{}
		fmt.Println("All files saved")
	}()

	for i := 0; i < b.maxConcurrentWorkers; i++ {
		if b.delayUntilInvokeNewWorker > 0 {
			time.Sleep(b.delayUntilInvokeNewWorker)
		}
		go b.save()
	}

	<-endChan
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
	b.wg.Add(1)
	defer b.wg.Done()
	for f := range b.chunksChan {
		err := b.chunkHandler.Save(f)
		if err != nil {
			b.errsChan <- err
		}
	}
}
