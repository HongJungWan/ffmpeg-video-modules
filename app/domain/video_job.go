package domain

import (
	"encoding/json"
	"time"
)

type VideoJob struct {
	ID             int                    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	VideoID        int                    `gorm:"not null" json:"videoId"`
	JobType        VideoJobType           `gorm:"type:ENUM('trim','concat');not null" json:"jobType"`
	Parameters     map[string]interface{} `gorm:"type:json;not null" json:"parameters"` // MySQL 5.7+ supports JSON
	Status         JobStatus              `gorm:"type:ENUM('pending','in_progress','completed','failed');default:'pending'" json:"status"`
	ResultFilePath string                 `gorm:"type:varchar(255)" json:"resultFilePath"`
	CreatedAt      time.Time              `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time              `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

type VideoJobType string

const (
	Trim   VideoJobType = "trim"
	Concat VideoJobType = "concat"
)

type JobStatus string

const (
	Pending    JobStatus = "pending"
	InProgress JobStatus = "in_progress"
	Completed  JobStatus = "completed"
	JobFailed  JobStatus = "failed"
)

func NewVideoJob(videoID int, jobType VideoJobType, parameters map[string]interface{}) *VideoJob {
	// map을 JSON 문자열로 변환
	parametersJSON, _ := json.Marshal(parameters)

	return &VideoJob{
		VideoID:    videoID,
		JobType:    jobType,
		Parameters: string(parametersJSON),
		Status:     Pending,
		CreatedAt:  time.Now(),
	}
}

func (vj *VideoJob) UpdateStatus(status JobStatus) {
	vj.Status = status
	vj.UpdatedAt = time.Now()
}
