package response

import (
	"time"
)

type VideoResponse struct {
	ID        int       `json:"id"`
	Filename  string    `json:"filename"`
	FilePath  string    `json:"filePath"`
	Duration  int       `json:"duration"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
