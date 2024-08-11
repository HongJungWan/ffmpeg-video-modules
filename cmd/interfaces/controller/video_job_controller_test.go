package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/HongJungWan/ffmpeg-video-modules/cmd/interfaces/controller"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/interfaces/dto/request"
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

func (m *MockVideoJobInteractor) ExecuteJobs() error {
	args := m.Called()
	return args.Error(0)
}

func TestTrimVideo(t *testing.T) {
	// Given
	videoID := 1
	reqBody := request.TrimVideoRequest{
		TrimStart: "00:00:05",
		TrimEnd:   "00:00:10",
	}
	jsonValue, _ := json.Marshal(reqBody)

	videoJobInteractor := new(MockVideoJobInteractor)
	videoJobInteractor.On("TrimVideo", videoID, reqBody.TrimStart, reqBody.TrimEnd).Return(1, nil)

	r := gin.Default()
	vjc := controller.NewVideoJobController(videoJobInteractor)
	r.POST("/video/:id/trim", vjc.TrimVideo)

	req, _ := http.NewRequest(http.MethodPost, "/video/"+strconv.Itoa(videoID)+"/trim", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// When
	r.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "jobID")
	videoJobInteractor.AssertCalled(t, "TrimVideo", videoID, reqBody.TrimStart, reqBody.TrimEnd)
}

func TestConcatVideos(t *testing.T) {
	// Given
	reqBody := request.ConcatVideosRequest{
		VideoIDs: []int{1, 2},
	}
	jsonValue, _ := json.Marshal(reqBody)

	videoJobInteractor := new(MockVideoJobInteractor)
	videoJobInteractor.On("ConcatVideos", reqBody.VideoIDs).Return(1, nil)

	r := gin.Default()
	vjc := controller.NewVideoJobController(videoJobInteractor)
	r.POST("/video/concat", vjc.ConcatVideos)

	req, _ := http.NewRequest(http.MethodPost, "/video/concat", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// When
	r.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "jobID")
	videoJobInteractor.AssertCalled(t, "ConcatVideos", reqBody.VideoIDs)
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
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "All jobs executed successfully")
	videoJobInteractor.AssertCalled(t, "ExecuteJobs")
}
