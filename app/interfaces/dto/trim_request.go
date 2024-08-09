package dto

type TrimVideoRequest struct {
	TrimStart string `json:"trimStart" binding:"required"`
	TrimEnd   string `json:"trimEnd" binding:"required"`
}
