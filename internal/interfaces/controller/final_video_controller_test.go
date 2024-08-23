package controller_test

import (
	"github.com/HongJungWan/ffmpeg-video-modules/internal/interfaces/controller"
	"github.com/HongJungWan/ffmpeg-video-modules/internal/interfaces/dto/response"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFinalVideoInteractor struct {
	mock.Mock
}

func (m *MockFinalVideoInteractor) GetFinalVideoDownloadLink(videoID int) (*response.VideoDownloadResponse, error) {
	args := m.Called(videoID)
	return args.Get(0).(*response.VideoDownloadResponse), args.Error(1)
}

func TestDownloadFinalVideo(t *testing.T) {
	// Given
	videoID := 1
	expectedResponse := &response.VideoDownloadResponse{
		ID:           videoID,
		Filename:     "example.mp4",
		FilePath:     "/videos/example.mp4",
		DownloadLink: "https://example.com/download",
	}

	finalVideoInteractor := new(MockFinalVideoInteractor)
	finalVideoInteractor.On("GetFinalVideoDownloadLink", videoID).Return(expectedResponse, nil)

	r := gin.Default()
	fvc := controller.NewFinalVideoController(finalVideoInteractor)
	r.GET("/video/:fid/download", fvc.DownloadFinalVideo)

	req, _ := http.NewRequest(http.MethodGet, "/video/"+strconv.Itoa(videoID)+"/download", nil)
	w := httptest.NewRecorder()

	// When
	r.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "https://example.com/download")
	finalVideoInteractor.AssertCalled(t, "GetFinalVideoDownloadLink", videoID)
}
