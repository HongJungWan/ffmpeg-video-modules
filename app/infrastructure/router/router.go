package router

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/infrastructure/configs"
	"github.com/HongJungWan/ffmpeg-video-modules/app/interfaces"
	"github.com/HongJungWan/ffmpeg-video-modules/app/usecases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(conf configs.Config) *gin.Engine {
	service := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true

	service.Use(cors.New(config))
	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ffmpeg video modules")
	})

	router := service.Group("/api")

	// Use Case 및 Controller 설정
	healthCheckInteractor := usecases.NewHealthCheckInteractor()
	healthCheckController := interfaces.NewHealthCheckController(healthCheckInteractor)
	router.GET("/health", healthCheckController.HealthCheck)

	return service
}
