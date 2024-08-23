package domain

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewVideoJob(t *testing.T) {
	// Given
	videoID := 1
	jobType := Trim
	parameters := map[string]interface{}{
		"startTime": "00:00:10",
		"endTime":   "00:01:00",
	}

	// When
	videoJob := NewVideoJob(videoID, jobType, parameters)

	// Then
	assert.Equal(t, videoID, videoJob.VideoID, "VideoID는 1이어야 합니다.")
	assert.Equal(t, jobType, videoJob.JobType, "JobType은 'trim'이어야 합니다.")
	assert.Equal(t, Pending, videoJob.Status, "초기 상태는 'pending'이어야 합니다.")

	expectedParameters, _ := json.Marshal(parameters)
	assert.Equal(t, string(expectedParameters), videoJob.Parameters, "Parameters는 JSON 문자열로 올바르게 변환되어야 합니다.")

	assert.WithinDuration(t, time.Now(), videoJob.CreatedAt, time.Second, "CreatedAt은 현재 시간과 거의 일치해야 합니다.")
}

func TestUpdateVideoJobStatus(t *testing.T) {
	// Given
	videoJob := NewVideoJob(1, Trim, map[string]interface{}{
		"startTime": "00:00:10",
		"endTime":   "00:01:00",
	})
	initialStatus := videoJob.Status

	// When
	videoJob.UpdateStatus(InProgress)

	// Then
	assert.Equal(t, InProgress, videoJob.Status, "VideoJob 상태는 'in_progress'로 업데이트되어야 합니다.")
	assert.NotEqual(t, initialStatus, videoJob.Status, "상태는 초기 상태와 달라야 합니다.")
	assert.WithinDuration(t, time.Now(), videoJob.UpdatedAt, time.Second, "UpdatedAt은 현재 시간과 거의 일치해야 합니다.")
}
