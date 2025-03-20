package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSS3ChunkHandler struct {
	ctx    context.Context
	client *s3.Client
	bucket string
	path   string
	Delay  time.Duration
}

func NewAWSS3ChunkHandler(ctx context.Context, delay time.Duration) *AWSS3ChunkHandler {
	bucket := os.Getenv("AWS_S3_BUCKET")
	if bucket == "" {
		panic("AWS_S3_BUCKET env var is required")
	}
	path := os.Getenv("AWS_S3_PATH")
	if path == "" {
		panic("AWS_S3_PATH env var is required")
	}
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	if path[0] == '/' {
		path = path[1:]
	}
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("sa-east-1"))
	if err != nil {
		panic(fmt.Errorf("unable to load SDK config, %v", err))
	}
	return &AWSS3ChunkHandler{
		ctx:    ctx,
		Delay:  delay,
		bucket: bucket,
		path:   path,
		client: s3.NewFromConfig(cfg),
	}
}

func (h AWSS3ChunkHandler) delayBetweenChunks() {
	time.Sleep(h.Delay)
}

func (h AWSS3ChunkHandler) Save(c *Chunk) error {
	h.delayBetweenChunks()

	start := time.Now()

	fmt.Printf("Saving: %s.csv...\n", c.name)
	_, err := h.client.PutObject(h.ctx, &s3.PutObjectInput{
		Bucket: aws.String(h.bucket),
		Key:    aws.String(fmt.Sprintf("input/%s/%s.csv", h.path, c.name)),
		Body:   bytes.NewReader(c.data),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Saved: %s.csv in %s\n", c.name, time.Since(start))
	return nil
}
