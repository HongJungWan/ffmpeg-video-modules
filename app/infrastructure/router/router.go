package router

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/app/infrastructure/configs"
	"github.com/HongJungWan/ffmpeg-video-modules/app/interfaces/controller"
	"github.com/HongJungWan/ffmpeg-video-modules/app/interfaces/repository"
	"github.com/HongJungWan/ffmpeg-video-modules/app/usecases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func NewRouter(conf configs.Config, db *gorm.DB) *gin.Engine {
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

	// 데이터베이스 마이그레이션
	db.Table("video").AutoMigrate(&domain.Video{})
	db.Table("video_job").AutoMigrate(&domain.VideoJob{})
	db.Table("final_video").AutoMigrate(&domain.FinalVideo{})

	// Health Check 관련 설정
	healthCheckInteractor := usecases.NewHealthCheckInteractor()
	healthCheckController := controller.NewHealthCheckController(healthCheckInteractor)

	// Video 관련 설정
	videoRepository := repository.NewVideoRepository(db)
	videoInteractor := usecases.NewVideoInteractor(videoRepository)
	videoController := controller.NewVideoController(videoInteractor)

	// VideoJob 관련 설정 (Trim, Concat, Execute)
	videoJobRepository := repository.NewVideoJobRepository(db)
	videoJobInteractor := usecases.NewVideoJobInteractor(videoJobRepository, videoRepository)
	videoJobController := controller.NewVideoJobController(videoJobInteractor)

	router := service.Group("/api")

	router.GET("/health", healthCheckController.HealthCheck)

	router.POST("/videos", videoController.UploadVideo)
	router.POST("/videos/:id/trim", videoJobController.TrimVideo)
	router.POST("/videos/concat", videoJobController.ConcatVideos)

	router.POST("/jobs/execute", videoJobController.ExecuteJobs)

	return service
}
