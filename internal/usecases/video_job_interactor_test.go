package usecases_test

import (
	domain2 "github.com/HongJungWan/ffmpeg-video-modules/internal/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/test/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/HongJungWan/ffmpeg-video-modules/cmd/usecases"
	"github.com/stretchr/testify/assert"
)

func TestTrimVideo(t *testing.T) {
	// Given
	videoID := 1
	mockVideoRepo := new(mocks.MockVideoRepository)
	mockVideoRepo.On("FindByID", videoID).Return(&domain2.Video{
		ID:       videoID,
		Filename: "example.mp4",
	}, nil)

	mockJobRepo := new(mocks.MockVideoJobRepository)
	// Save 호출 시 job.ID를 1로 설정
	mockJobRepo.On("Save", mock.Anything).Run(func(args mock.Arguments) {
		job := args.Get(0).(*domain2.VideoJob)
		job.ID = 1 // 임의의 ID 설정
	}).Return(nil)

	interactor := usecases.NewVideoJobInteractor(mockJobRepo, mockVideoRepo, nil)

	// When
	jobID, err := interactor.TrimVideo(videoID, "00:00:05", "00:00:10")

	// Then
	assert.NoError(t, err)
	assert.NotZero(t, jobID) // jobID가 0이 아님을 확인
	mockVideoRepo.AssertCalled(t, "FindByID", videoID)
	mockJobRepo.AssertCalled(t, "Save", mock.Anything)
}

func TestConcatVideos(t *testing.T) {
	// Given
	videoIDs := []int{1, 2}
	mockVideoRepo := new(mocks.MockVideoRepository)
	mockVideoRepo.On("FindByID", 1).Return(&domain2.Video{
		ID:       1,
		Filename: "example1.mp4",
	}, nil)
	mockVideoRepo.On("FindByID", 2).Return(&domain2.Video{
		ID:       2,
		Filename: "example2.mp4",
	}, nil)

	mockJobRepo := new(mocks.MockVideoJobRepository)
	// Save 호출 시 job.ID를 1로 설정
	mockJobRepo.On("Save", mock.Anything).Run(func(args mock.Arguments) {
		job := args.Get(0).(*domain2.VideoJob)
		job.ID = 1 // 임의의 ID 설정
	}).Return(nil)

	interactor := usecases.NewVideoJobInteractor(mockJobRepo, mockVideoRepo, nil)

	// When
	jobID, err := interactor.ConcatVideos(videoIDs)

	// Then
	assert.NoError(t, err)
	assert.NotZero(t, jobID) // jobID가 0이 아님을 확인
	mockVideoRepo.AssertCalled(t, "FindByID", 1)
	mockVideoRepo.AssertCalled(t, "FindByID", 2)
	mockJobRepo.AssertCalled(t, "Save", mock.Anything)
}
