package controller_test

import (
	"github.com/HongJungWan/ffmpeg-video-modules/internal/interfaces/controller"
	response2 "github.com/HongJungWan/ffmpeg-video-modules/internal/interfaces/dto/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVideoInteractor struct {
	mock.Mock
}

func (m *MockVideoInteractor) GetVideoDetails() ([]response2.VideoDetailResponse, error) {
	args := m.Called()
	return args.Get(0).([]response2.VideoDetailResponse), args.Error(1)
}

func (m *MockVideoInteractor) HandleVideoUpload(ctx *gin.Context) ([]response2.VideoResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]response2.VideoResponse), args.Error(1)
}

func TestGetVideoDetails(t *testing.T) {
	// Given
	expectedResponse := []response2.VideoDetailResponse{
		{
			ID:        1,
			Filename:  "example.mp4",
			FilePath:  "/videos/example.mp4",
			Duration:  120,
			Status:    "completed",
			CreatedAt: "2023-08-11T12:00:00Z",
			UpdatedAt: "2023-08-11T12:30:00Z",
		},
	}

	videoInteractor := new(MockVideoInteractor)
	videoInteractor.On("GetVideoDetails").Return(expectedResponse, nil)

	r := gin.Default()
	vdc := controller.NewVideoController(videoInteractor)
	r.GET("/video/details", vdc.GetVideoDetails)

	req, _ := http.NewRequest(http.MethodGet, "/video/details", nil)
	w := httptest.NewRecorder()

	// When
	r.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "example.mp4")
	videoInteractor.AssertCalled(t, "GetVideoDetails")
}

func TestUploadVideo(t *testing.T) {
	// Given
	expectedResponse := []response2.VideoResponse{
		{
			ID:       1,
			Filename: "uploaded_video.mp4",
			FilePath: "/uploads/uploaded_video.mp4",
			Duration: 100,
			Status:   "uploaded",
		},
	}

	videoInteractor := new(MockVideoInteractor)
	videoInteractor.On("HandleVideoUpload", mock.Anything).Return(expectedResponse, nil)

	r := gin.Default()
	vc := controller.NewVideoController(videoInteractor)
	r.POST("/video/upload", vc.UploadVideo)

	req, _ := http.NewRequest(http.MethodPost, "/video/upload", nil)
	w := httptest.NewRecorder()

	// When
	r.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "success")
	videoInteractor.AssertCalled(t, "HandleVideoUpload", mock.Anything)
}
