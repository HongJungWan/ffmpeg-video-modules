package request

type ConcatVideosRequest struct {
	VideoIDs []int `json:"videoIds" binding:"required"`
}
