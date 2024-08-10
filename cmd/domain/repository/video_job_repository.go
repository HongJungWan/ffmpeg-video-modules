package repository

import "github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"

type VideoJobRepository interface {
	Save(job *domain.VideoJob) error
	FindByID(id int) (*domain.VideoJob, error)
	FindPendingJobs() ([]*domain.VideoJob, error)
	UpdateStatus(job *domain.VideoJob) error
	FindByVideoIDAndType(videoID int, jobType domain.VideoJobType) ([]*domain.VideoJob, error)
}
