package controller

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoController struct {
	videoInteractor usecases.VideoInteractor
}

func NewVideoController(videoInteractor usecases.VideoInteractor) *VideoController {
	return &VideoController{videoInteractor: videoInteractor}
}

func (vc *VideoController) UploadVideo(ctx *gin.Context) {
	err := vc.videoInteractor.HandleVideoUpload(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
