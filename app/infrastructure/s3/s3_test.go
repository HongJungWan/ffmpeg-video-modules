package s3

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/stretchr/testify/assert"
)

type MockS3Client struct {
	PutObjectFunc     func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObjectFunc     func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	ListObjectsV2Func func(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	HeadObjectFunc    func(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error)
}

func (m *MockS3Client) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	return m.PutObjectFunc(ctx, params, optFns...)
}

func (m *MockS3Client) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return m.GetObjectFunc(ctx, params, optFns...)
}

func (m *MockS3Client) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	return m.ListObjectsV2Func(ctx, params, optFns...)
}

func (m *MockS3Client) HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	return m.HeadObjectFunc(ctx, params, optFns...)
}

type MockMultipartFile struct {
	*bytes.Reader
}

func (m *MockMultipartFile) Header() *multipart.FileHeader {
	return &multipart.FileHeader{}
}

func (m *MockMultipartFile) Close() error {
	return nil
}

func setupMockClient() (*Client, *MockS3Client) {
	mockClient := &MockS3Client{}
	client := &Client{s3: mockClient}
	return client, mockClient
}

func TestDownloadFile(t *testing.T) {
	// Given
	client, mockClient := setupMockClient()

	mockClient.GetObjectFunc = func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
		return &s3.GetObjectOutput{
			Body: io.NopCloser(bytes.NewReader([]byte("video content"))),
		}, nil
	}

	bucketName := "test-bucket"
	key := "test-video.mp4"

	// When
	body, err := client.DownloadFile(bucketName, key)
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, body)
	assert.Equal(t, "video content", buf.String())
}

func TestListFiles(t *testing.T) {
	// Given
	client, mockClient := setupMockClient()

	mockClient.ListObjectsV2Func = func(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
		return &s3.ListObjectsV2Output{
			Contents: []types.Object{
				{Key: aws.String("video1.mp4")},
				{Key: aws.String("video2.mp4")},
			},
		}, nil
	}

	bucketName := "test-bucket"

	// When
	result, err := client.ListFiles(bucketName)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Contents, 2)
	assert.Equal(t, "video1.mp4", *result.Contents[0].Key)
	assert.Equal(t, "video2.mp4", *result.Contents[1].Key)
}

func TestGetFileMetadata(t *testing.T) {
	// Given
	client, mockClient := setupMockClient()

	mockClient.HeadObjectFunc = func(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
		return &s3.HeadObjectOutput{
			ContentLength: aws.Int64(12345),
			ContentType:   aws.String("video/mp4"),
		}, nil
	}

	bucketName := "test-bucket"
	key := "test-video.mp4"

	// When
	result, err := client.GetFileMetadata(bucketName, key)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "video/mp4", *result.ContentType)
}
