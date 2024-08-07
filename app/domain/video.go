package domain

import (
	"time"
)

type Video struct {
	ID        int         `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Filename  string      `gorm:"type:varchar(255);not null" json:"filename"`
	FilePath  string      `gorm:"type:varchar(255);not null" json:"filePath"`
	Status    VideoStatus `gorm:"type:ENUM('uploaded','processed','failed');default:'uploaded'" json:"status"`
	CreatedAt time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

type VideoStatus string

const (
	Uploaded  VideoStatus = "uploaded"
	Processed VideoStatus = "processed"
	Failed    VideoStatus = "failed"
)

func NewVideo(filename, filePath string) *Video {
	return &Video{
		Filename:  filename,
		FilePath:  filePath,
		Status:    Uploaded,
		CreatedAt: time.Now(),
	}
}

func (v *Video) UpdateStatus(status VideoStatus) {
	v.Status = status
	v.UpdatedAt = time.Now()
}
