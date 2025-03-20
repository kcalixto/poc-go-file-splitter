package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type LocalFileSystemHandler struct{}

func (h LocalFileSystemHandler) Save(c *Chunk) error {
	fmt.Printf("Saving %s...\n", c.name)
	time.Sleep(1 * time.Second)
	file, err := os.Create(fmt.Sprintf("out/%s.csv", c.name))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(c.data))
	if err != nil {
		return err
	}

	return nil
}
