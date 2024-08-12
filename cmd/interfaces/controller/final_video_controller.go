package controller

import (
	"net/http"
	"strconv"

	"github.com/HongJungWan/ffmpeg-video-modules/cmd/usecases"
	"github.com/gin-gonic/gin"
)

type FinalVideoController struct {
	finalVideoInteractor usecases.FinalVideoInteractor
}

func NewFinalVideoController(finalVideoInteractor usecases.FinalVideoInteractor) *FinalVideoController {
	return &FinalVideoController{finalVideoInteractor: finalVideoInteractor}
}

// DownloadFinalVideo godoc
// @Summary      Downloads the final video
// @Description  Retrieves the download link for the final video based on the video ID
// @Tags         video
// @Accept       json
// @Produce      json
// @Param        fid   path      int  true  "Final Video ID"
// @Success      200 {object} response.VideoDownloadResponse "Download link for the final video"
// @Failure      400 {object} map[string]interface{} "Invalid video ID"
// @Failure      500 {object} map[string]interface{} "Internal Server Error"
// @Router       /videos/{fid}/download [get]
func (fvc *FinalVideoController) DownloadFinalVideo(ctx *gin.Context) {
	videoID, err := strconv.Atoi(ctx.Param("fid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	response, err := fvc.finalVideoInteractor.GetFinalVideoDownloadLink(videoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
