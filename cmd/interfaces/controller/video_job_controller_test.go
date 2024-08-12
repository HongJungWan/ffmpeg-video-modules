package controller_test

import (
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/interfaces/dto/response"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/usecases"
	"github.com/HongJungWan/ffmpeg-video-modules/test/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HongJungWan/ffmpeg-video-modules/cmd/interfaces/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVideoJobInteractor struct {
	mock.Mock
}

func (m *MockVideoJobInteractor) TrimVideo(videoID int, trimStart string, trimEnd string) (int, error) {
	args := m.Called(videoID, trimStart, trimEnd)
	return args.Int(0), args.Error(1)
}

func (m *MockVideoJobInteractor) ConcatVideos(videoIDs []int) (int, error) {
	args := m.Called(videoIDs)
	return args.Int(0), args.Error(1)
}

func (m *MockVideoJobInteractor) ExecuteJobs(jobIDs []int) ([]response.JobIDResponse, error) {
	args := m.Called(jobIDs)
	return args.Get(0).([]response.JobIDResponse), args.Error(1)
}

func TestTrimVideo(t *testing.T) {
	// Given
	videoID := 1
	mockVideoRepo := new(mocks.MockVideoRepository)
	mockVideoRepo.On("FindByID", videoID).Return(&domain.Video{
		ID:       videoID,
		Filename: "example.mp4",
	}, nil)

	mockJobRepo := new(mocks.MockVideoJobRepository)
	mockJobRepo.On("Save", mock.Anything).Run(func(args mock.Arguments) {
		job := args.Get(0).(*domain.VideoJob)
		job.ID = 1 // 임의의 ID 설정
	}).Return(nil)

	interactor := usecases.NewVideoJobInteractor(mockJobRepo, mockVideoRepo, nil)

	// When
	jobID, err := interactor.TrimVideo(videoID, "00:00:05", "00:00:10")

	// Then
	assert.NoError(t, err)
	assert.NotZero(t, jobID)
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
	mockJobRepo.On("Save", mock.Anything).Run(func(args mock.Arguments) {
		job := args.Get(0).(*domain.VideoJob)
		job.ID = 1 // 임의의 ID 설정
	}).Return(nil)

	interactor := usecases.NewVideoJobInteractor(mockJobRepo, mockVideoRepo, nil)

	// When
	jobID, err := interactor.ConcatVideos(videoIDs)

	// Then
	assert.NoError(t, err)
	assert.NotZero(t, jobID)
	mockVideoRepo.AssertCalled(t, "FindByID", 1)
	mockVideoRepo.AssertCalled(t, "FindByID", 2)
	mockJobRepo.AssertCalled(t, "Save", mock.Anything)
}

func TestExecuteJobs(t *testing.T) {
	// Given
	videoJobInteractor := new(MockVideoJobInteractor)
	videoJobInteractor.On("ExecuteJobs").Return(nil)

	r := gin.Default()
	vjc := controller.NewVideoJobController(videoJobInteractor)
	r.POST("/video/jobs/execute", vjc.ExecuteJobs)

	req, _ := http.NewRequest(http.MethodPost, "/video/jobs/execute", nil)
	w := httptest.NewRecorder()

	// When
	r.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "All jobs executed successfully")
	videoJobInteractor.AssertCalled(t, "ExecuteJobs")
}
