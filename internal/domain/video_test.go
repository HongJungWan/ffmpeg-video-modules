package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewVideo(t *testing.T) {
	// Given
	filename := "example.mp4"
	filePath := "/videos/example.mp4"
	duration := 120

	// When
	video := NewVideo(filename, filePath, duration)

	// Then
	assert.Equal(t, filename, video.Filename, "비디오 파일명은 'example.mp4'여야 합니다.")
	assert.Equal(t, filePath, video.FilePath, "비디오 파일 경로는 '/videos/example.mp4'여야 합니다.")
	assert.Equal(t, duration, video.Duration, "비디오 길이는 120이어야 합니다.")
	assert.Equal(t, Uploaded, video.Status, "초기 상태는 'uploaded'여야 합니다.")
	assert.WithinDuration(t, time.Now(), video.CreatedAt, time.Second, "CreatedAt은 현재 시간과 거의 일치해야 합니다.")
}

func TestUpdateStatus(t *testing.T) {
	// Given
	video := NewVideo("example.mp4", "/videos/example.mp4", 120)
	initialStatus := video.Status

	// When
	video.UpdateStatus(Processed)

	// Then
	assert.Equal(t, Processed, video.Status, "비디오 상태는 'processed'로 업데이트되어야 합니다.")
	assert.NotEqual(t, initialStatus, video.Status, "상태는 초기 상태와 달라야 합니다.")
	assert.WithinDuration(t, time.Now(), video.UpdatedAt, time.Second, "UpdatedAt은 현재 시간과 거의 일치해야 합니다.")
}
