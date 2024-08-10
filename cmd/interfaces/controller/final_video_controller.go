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
