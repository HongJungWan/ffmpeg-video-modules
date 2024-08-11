package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFinalVideo(t *testing.T) {
	// Given
	originalVideoID := 1
	filename := "final_example.mp4"
	filePath := "/videos/final_example.mp4"
	duration := 180

	// When
	finalVideo := NewFinalVideo(originalVideoID, filename, filePath, duration)

	// Then
	assert.Equal(t, originalVideoID, finalVideo.OriginalVideoID, "OriginalVideoID는 1이어야 합니다.")
	assert.Equal(t, filename, finalVideo.Filename, "Filename은 'final_example.mp4'이어야 합니다.")
	assert.Equal(t, filePath, finalVideo.FilePath, "FilePath는 '/videos/final_example.mp4'이어야 합니다.")
	assert.Equal(t, duration, finalVideo.Duration, "Duration은 180이어야 합니다.")
	assert.Equal(t, Processed, finalVideo.Status, "초기 상태는 'processed'여야 합니다.")
	assert.WithinDuration(t, time.Now(), finalVideo.CreatedAt, time.Second, "CreatedAt은 현재 시간과 거의 일치해야 합니다.")
}

func TestUpdateFinalVideoStatus(t *testing.T) {
	// Given
	finalVideo := NewFinalVideo(1, "final_example.mp4", "/videos/final_example.mp4", 180)
	initialStatus := finalVideo.Status

	// When
	finalVideo.UpdateStatus(Failed)

	// Then
	assert.Equal(t, Failed, finalVideo.Status, "FinalVideo 상태는 'failed'로 업데이트되어야 합니다.")
	assert.NotEqual(t, initialStatus, finalVideo.Status, "상태는 초기 상태와 달라야 합니다.")
	assert.WithinDuration(t, time.Now(), finalVideo.UpdatedAt, time.Second, "UpdatedAt은 현재 시간과 거의 일치해야 합니다.")
}
