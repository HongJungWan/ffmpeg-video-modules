package controller

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/interfaces/dto/request"
	"net/http"
	"strconv"

	"github.com/HongJungWan/ffmpeg-video-modules/app/usecases"
	"github.com/gin-gonic/gin"
)

type VideoJobController struct {
	videoJobInteractor usecases.VideoJobInteractor
}

func NewVideoJobController(videoJobInteractor usecases.VideoJobInteractor) *VideoJobController {
	return &VideoJobController{videoJobInteractor: videoJobInteractor}
}

func (vjc *VideoJobController) TrimVideo(ctx *gin.Context) {
	videoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	var req request.TrimVideoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobID, err := vjc.videoJobInteractor.TrimVideo(videoID, req.TrimStart, req.TrimEnd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jobID": jobID})
}

func (vjc *VideoJobController) ConcatVideos(ctx *gin.Context) {
	var req request.ConcatVideosRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobID, err := vjc.videoJobInteractor.ConcatVideos(req.VideoIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jobID": jobID})
}

func (vjc *VideoJobController) ExecuteJobs(ctx *gin.Context) {
	err := vjc.videoJobInteractor.ExecuteJobs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "All jobs executed successfully"})
}
