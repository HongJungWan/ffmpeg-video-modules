package controller

import (
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/interfaces/dto/request"
	"net/http"
	"strconv"

	"github.com/HongJungWan/ffmpeg-video-modules/cmd/usecases"
	"github.com/gin-gonic/gin"
)

type VideoJobController struct {
	videoJobInteractor usecases.VideoJobInteractor
}

func NewVideoJobController(videoJobInteractor usecases.VideoJobInteractor) *VideoJobController {
	return &VideoJobController{videoJobInteractor: videoJobInteractor}
}

// TrimVideo godoc
// @Summary      Trims a video
// @Description  Trims the specified video to the given start and end times
// @Tags         video
// @Accept       json
// @Produce      json
// @Param        id   path      int                     true  "Video ID"
// @Param        body body      request.TrimVideoRequest true  "Trim start and end times"
// @Success      202 {object} map[string]int            "Job ID of the trimming task"
// @Failure      400 {object} map[string]interface{}    "Invalid video ID or bad request"
// @Failure      500 {object} map[string]interface{}    "Internal Server Error"
// @Router       /videos/{id}/trim [post]
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

	ctx.JSON(http.StatusAccepted, gin.H{"jobID": jobID})
}

// ConcatVideos godoc
// @Summary      Concatenates multiple videos
// @Description  Concatenates the specified videos into a single video
// @Tags         video
// @Accept       json
// @Produce      json
// @Param        body body      request.ConcatVideosRequest true  "List of video IDs to concatenate"
// @Success      202 {object} map[string]int               "Job ID of the concatenation task"
// @Failure      400 {object} map[string]interface{}       "Bad request"
// @Failure      500 {object} map[string]interface{}       "Internal Server Error"
// @Router       /videos/concat [post]
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

	ctx.JSON(http.StatusAccepted, gin.H{"jobID": jobID})
}

// ExecuteJobs godoc
// @Summary      Executes pending video jobs
// @Description  Executes all pending video jobs, such as trimming or concatenation
// @Tags         job execute
// @Produce      json
// @Success      201 {object} map[string]string          "Status of job execution"
// @Failure      500 {object} map[string]interface{}     "Internal Server Error"
// @Router       /jobs/execute [post]
func (vjc *VideoJobController) ExecuteJobs(ctx *gin.Context) {
	err := vjc.videoJobInteractor.ExecuteJobs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "All jobs executed successfully"})
}
