package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/google/uuid"
)

type LocalFileSystemHandler struct{}

func (h LocalFileSystemHandler) Save(c *Chunk) error {
	name := uuid.New().String()
	fmt.Println("Saving file... ", name)
	time.Sleep(1 * time.Second)
	file, err := os.Create(fmt.Sprintf("out/%s.csv", name))
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
