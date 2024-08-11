package usecases_test

import (
	"github.com/HongJungWan/ffmpeg-video-modules/test/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/usecases"
	"github.com/stretchr/testify/assert"
)

func TestTrimVideo(t *testing.T) {
	// Given
	videoID := 1
	mockVideoRepo := new(mocks.MockVideoRepository)
	mockVideoRepo.On("FindByID", videoID).Return(&domain.Video{
		ID:       videoID,
		Filename: "example.mp4",
	}, nil)

	mockJobRepo := new(mocks.MockVideoJobRepository)
	mockJobRepo.On("Save", mock.Anything).Return(nil)

	interactor := usecases.NewVideoJobInteractor(mockJobRepo, mockVideoRepo, nil)

	// When
	_, err := interactor.TrimVideo(videoID, "00:00:05", "00:00:10")

	// Then
	assert.NoError(t, err)
	mockVideoRepo.AssertCalled(t, "FindByID", videoID)
	mockJobRepo.AssertCalled(t, "Save", mock.Anything)
}

func TestConcatVideos(t *testing.T) {
	// Given
	videoIDs := []int{1, 2}
	mockVideoRepo := new(mocks.MockVideoRepository)
	mockVideoRepo.On("FindByID", 1).Return(&domain.Video{
		ID:       1,
		Filename: "example1.mp4",
	}, nil)
	mockVideoRepo.On("FindByID", 2).Return(&domain.Video{
		ID:       2,
		Filename: "example2.mp4",
	}, nil)

	mockJobRepo := new(mocks.MockVideoJobRepository)
	mockJobRepo.On("Save", mock.Anything).Return(nil)

	interactor := usecases.NewVideoJobInteractor(mockJobRepo, mockVideoRepo, nil)

	// When
	_, err := interactor.ConcatVideos(videoIDs)

	// Then
	assert.NoError(t, err)
	mockVideoRepo.AssertCalled(t, "FindByID", 1)
	mockVideoRepo.AssertCalled(t, "FindByID", 2)
	mockJobRepo.AssertCalled(t, "Save", mock.Anything)
}
