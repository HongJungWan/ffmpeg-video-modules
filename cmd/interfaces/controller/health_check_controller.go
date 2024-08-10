package controller

import (
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckController struct {
	interactor usecases.HealthCheckInteractor
}

func NewHealthCheckController(interactor usecases.HealthCheckInteractor) *HealthCheckController {
	return &HealthCheckController{interactor}
}

func (h *HealthCheckController) HealthCheck(c *gin.Context) {
	result := h.interactor.PerformHealthCheck()
	c.JSON(http.StatusOK, result)
}
