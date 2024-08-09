package dto

import (
	"time"
)

type VideoResponse struct {
	ID        int       `json:"id"`
	Filename  string    `json:"filename"`
	FilePath  string    `json:"filePath"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
