package controller_test

import (
	"bytes"
	"encoding/json"
	domain2 "github.com/HongJungWan/ffmpeg-video-modules/internal/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/internal/interfaces/controller"
	"github.com/HongJungWan/ffmpeg-video-modules/internal/interfaces/dto/request"
	"github.com/HongJungWan/ffmpeg-video-modules/internal/interfaces/dto/response"
	"github.com/HongJungWan/ffmpeg-video-modules/internal/usecases"
	"github.com/HongJungWan/ffmpeg-video-modules/test/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

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
	mockVideoRepo.On("FindByID", videoID).Return(&domain2.Video{
		ID:       videoID,
		Filename: "example.mp4",
	}, nil)

	mockJobRepo := new(mocks.MockVideoJobRepository)
	mockJobRepo.On("Save", mock.Anything).Run(func(args mock.Arguments) {
		job := args.Get(0).(*domain2.VideoJob)
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
	mockVideoRepo.On("FindByID", 1).Return(&domain2.Video{
		ID:       1,
		Filename: "example1.mp4",
	}, nil)
	mockVideoRepo.On("FindByID", 2).Return(&domain2.Video{
		ID:       2,
		Filename: "example2.mp4",
	}, nil)

	mockJobRepo := new(mocks.MockVideoJobRepository)
	mockJobRepo.On("Save", mock.Anything).Run(func(args mock.Arguments) {
		job := args.Get(0).(*domain2.VideoJob)
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
	jobIDs := []int{1, 2} // 테스트를 위한 임의의 job IDs
	videoJobInteractor.On("ExecuteJobs", jobIDs).Return([]response.JobIDResponse{
		{JobID: 1},
		{JobID: 2},
	}, nil)

	r := gin.Default()
	vjc := controller.NewVideoJobController(videoJobInteractor)
	r.POST("/video/jobs/execute", vjc.ExecuteJobs)

	// 요청 본문에 포함할 job IDs JSON
	reqBody := request.ExecuteJobsRequest{JobIDs: jobIDs}
	jsonValue, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(http.MethodPost, "/video/jobs/execute", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// When
	r.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "jobID")
	videoJobInteractor.AssertCalled(t, "ExecuteJobs", jobIDs)
}
