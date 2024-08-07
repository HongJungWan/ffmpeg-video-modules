package main

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/infrastructure/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockRouter() http.Handler {
	return router.NewRouter(conf)
}

func TestStartServer(t *testing.T) {
	// Given
	router := mockRouter()

	// When
	req, err := http.NewRequest("GET", "/api/health", nil)
	assert.NoError(t, err, "요청 생성 중 오류가 발생했다.")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Then
	assert.Equal(t, http.StatusOK, rr.Code, "서버는 HTTP 200 상태 코드를 반환해야 한다.")
}
