package usecases_test

import (
	"github.com/HongJungWan/ffmpeg-video-modules/test/mocks"
	"testing"

	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/usecases"
	"github.com/stretchr/testify/assert"
)

func TestGetFinalVideoDownloadLink(t *testing.T) {
	// Given
	videoID := 1
	expectedFinalVideo := &domain.FinalVideo{
		ID:       videoID,
		Filename: "final_video.mp4",
		FilePath: "videos/final_video.mp4",
	}

	mockRepo := new(mocks.MockFinalVideoRepository)
	mockRepo.On("FindFinalVideoByID", videoID).Return(expectedFinalVideo, nil)

	interactor := usecases.NewFinalVideoInteractor(mockRepo)

	// When
	result, err := interactor.GetFinalVideoDownloadLink(videoID)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "final_video.mp4", result.Filename)
	assert.Equal(t, "http://localhost:3031/videos/final_video.mp4", result.DownloadLink)
	mockRepo.AssertCalled(t, "FindFinalVideoByID", videoID)
}
