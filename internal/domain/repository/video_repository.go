package repository

import (
	"github.com/HongJungWan/ffmpeg-video-modules/internal/domain"
)

type VideoRepository interface {
	Save(video *domain.Video) error
	FindByID(videoID int) (*domain.Video, error)
	FindAll() ([]domain.Video, error)
}
