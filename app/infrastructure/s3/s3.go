package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

type S3Client interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error)
}

type Client struct {
	s3 S3Client
}

func NewClient() (*Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("SDK 설정을 불러올 수 없습니다: %v", err)
	}

	return &Client{s3: s3.NewFromConfig(cfg)}, nil
}

func (c *Client) UploadFile(bucketName, key string, file multipart.File) (*s3.PutObjectOutput, error) {
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		return nil, err
	}

	mimeType := http.DetectContentType(buffer.Bytes())
	log.Printf("감지된 MIME 타입: %s", mimeType)

	result, err := c.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(buffer.Bytes()),
		ContentType: aws.String(mimeType),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) DownloadFile(bucketName, key string) (io.ReadCloser, error) {
	result, err := c.s3.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return result.Body, nil
}

func (c *Client) ListFiles(bucketName string) (*s3.ListObjectsV2Output, error) {
	result, err := c.s3.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetFileMetadata(bucketName, key string) (*s3.HeadObjectOutput, error) {
	result, err := c.s3.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
