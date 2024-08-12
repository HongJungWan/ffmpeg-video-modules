package router

import (
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/infrastructure/configs"
	repository_impl "github.com/HongJungWan/ffmpeg-video-modules/cmd/infrastructure/repository"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/interfaces/controller"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/usecases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	service.Static("/downloads", "./downloads")

	// Health Check 관련 설정
	healthCheckInteractor := usecases.NewHealthCheckInteractor()
	healthCheckController := controller.NewHealthCheckController(healthCheckInteractor)

	// 비즈니스 로직 관련 설정
	videoRepository := repository_impl.NewVideoRepository(db)
	videoJobRepository := repository_impl.NewVideoJobRepository(db)
	finalVideoRepository := repository_impl.NewFinalVideoRepository(db)

	videoInteractor := usecases.NewVideoInteractor(videoRepository, videoJobRepository, finalVideoRepository)
	videoJobInteractor := usecases.NewVideoJobInteractor(videoJobRepository, videoRepository, finalVideoRepository)
	finalVideoInteractor := usecases.NewFinalVideoInteractor(finalVideoRepository)

	videoController := controller.NewVideoController(videoInteractor)
	videoJobController := controller.NewVideoJobController(videoJobInteractor)
	finalVideoController := controller.NewFinalVideoController(finalVideoInteractor)

	router := service.Group("/api")

	// API 라우트 설정
	router.GET("/health", healthCheckController.HealthCheck)

	router.POST("/videos", videoController.UploadVideo)
	router.GET("/videos", videoController.GetVideoDetails)
	router.POST("/videos/:id/trim", videoJobController.TrimVideo)
	router.POST("/videos/concat", videoJobController.ConcatVideos)
	router.GET("/videos/:fid/download", finalVideoController.DownloadFinalVideo)

	router.POST("/jobs/execute", videoJobController.ExecuteJobs)

	// Swagger UI 라우트 추가
	service.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return service
}
