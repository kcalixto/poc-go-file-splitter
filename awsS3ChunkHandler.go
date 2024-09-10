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
	"github.com/google/uuid"
)

type AWSS3ChunkHandler struct {
	client *s3.Client
}

func (h *AWSS3ChunkHandler) Save(c *Chunk) error {
	name := uuid.New().String()
	if h.client == nil {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
		if err != nil {
			panic(fmt.Errorf("unable to load SDK config, %v", err))
		}
		h.client = s3.NewFromConfig(cfg)
	}

	bucket := os.Getenv("AWS_S3_BUCKET")
	if bucket == "" {
		panic("AWS_S3_BUCKET env var is required")
	}

	ctx := context.Background()
	start := time.Now()
	fmt.Printf("Saving: %s.csv...\n", name)
	_, err := h.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fmt.Sprintf("input/%s.csv", name)),
		Body:   bytes.NewReader(c.data),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Saved: %s.csv after %s\n", name, time.Since(start))
	return nil
}
