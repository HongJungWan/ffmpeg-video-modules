package controller

import (
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoController struct {
	videoInteractor usecases.VideoInteractor
}

func NewVideoController(videoInteractor usecases.VideoInteractor) *VideoController {
	return &VideoController{videoInteractor: videoInteractor}
}

// GetVideoDetails godoc
// @Summary      Retrieves details of all videos
// @Description  Fetches a list of all video details including trimming, concatenation jobs, and final video information
// @Tags         video
// @Accept       json
// @Produce      json
// @Success      200 {array} response.VideoDetailResponse "List of video details"
// @Failure      500 {object} map[string]interface{} "Internal Server Error"
// @Router       /videos [get]
func (vdc *VideoController) GetVideoDetails(ctx *gin.Context) {
	videoDetails, err := vdc.videoInteractor.GetVideoDetails()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, videoDetails)
}

// UploadVideo godoc
// @Summary      Uploads video files
// @Description  Handles the upload of multiple video files, saves them, and returns their details
// @Tags         video
// @Accept       multipart/form-data
// @Produce      json
// @Param        files  formData  file  true  "Video files to upload" collectionFormat=multi
// @Success      201 {array} response.VideoResponse "Uploaded video details"
// @Failure      500 {object} map[string]interface{} "Internal Server Error"
// @Router       /videos [post]
func (vc *VideoController) UploadVideo(ctx *gin.Context) {
	videoResponses, err := vc.videoInteractor.HandleVideoUpload(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"videos": videoResponses,
		"status": "success",
	})
}
