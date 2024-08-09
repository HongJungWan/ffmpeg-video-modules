package domain

import (
	"time"
)

type FinalVideo struct {
	ID              int         `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	OriginalVideoID int         `gorm:"not null" json:"originalVideoId"`
	Filename        string      `gorm:"type:varchar(255);not null" json:"filename"`
	FilePath        string      `gorm:"type:varchar(255);not null" json:"filePath"`
	Duration        int         `gorm:"type:int;not null" json:"duration"`
	Status          VideoStatus `gorm:"type:ENUM('processed','failed');default:'processed'" json:"status"`
	CreatedAt       time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt       time.Time   `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func NewFinalVideo(originalVideoID int, filename, filePath string, duration int) *FinalVideo {
	return &FinalVideo{
		OriginalVideoID: originalVideoID,
		Filename:        filename,
		FilePath:        filePath,
		Duration:        duration,
		Status:          Processed,
		CreatedAt:       time.Now(),
	}
}

func (fv *FinalVideo) UpdateStatus(status VideoStatus) {
	fv.Status = status
	fv.UpdatedAt = time.Now()
}
