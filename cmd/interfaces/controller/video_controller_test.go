package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HongJungWan/ffmpeg-video-modules/cmd/interfaces/controller"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/interfaces/dto/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVideoInteractor struct {
	mock.Mock
}

func (m *MockVideoInteractor) GetVideoDetails() ([]response.VideoDetailResponse, error) {
	args := m.Called()
	return args.Get(0).([]response.VideoDetailResponse), args.Error(1)
}

func (m *MockVideoInteractor) HandleVideoUpload(ctx *gin.Context) ([]response.VideoResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]response.VideoResponse), args.Error(1)
}

func TestGetVideoDetails(t *testing.T) {
	// Given
	expectedResponse := []response.VideoDetailResponse{
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
	expectedResponse := []response.VideoResponse{
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
