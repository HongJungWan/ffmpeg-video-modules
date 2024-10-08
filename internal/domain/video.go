package domain

import (
	"time"
)

type Video struct {
	ID        int         `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Filename  string      `gorm:"type:varchar(255);not null" json:"filename"`
	FilePath  string      `gorm:"type:varchar(255);not null" json:"filePath"`
	Duration  int         `gorm:"type:int;not null" json:"duration"`
	Status    VideoStatus `gorm:"type:ENUM('uploaded','processed','failed','finished');default:'uploaded'" json:"status"`
	CreatedAt time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

type VideoStatus string

const (
	Uploaded  VideoStatus = "uploaded"
	Processed VideoStatus = "processed"
	Failed    VideoStatus = "failed"
	FINISHED  VideoStatus = "finished"
)

func NewVideo(filename, filePath string, duration int) *Video {
	return &Video{
		Filename:  filename,
		FilePath:  filePath,
		Duration:  duration,
		Status:    Uploaded,
		CreatedAt: time.Now(),
	}
}

func (v *Video) UpdateStatus(status VideoStatus) {
	v.Status = status
	v.UpdatedAt = time.Now()
}
