package dto

type ConcatVideosRequest struct {
	VideoIDs []int `json:"videoIds" binding:"required"`
}
