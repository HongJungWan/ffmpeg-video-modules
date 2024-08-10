package repository

import "github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"

type FinalVideoRepository interface {
	SaveFinalVideo(finalVideo *domain.FinalVideo) error
	FindFinalVideoByID(id int) (*domain.FinalVideo, error)
	FindFinalVideoByOriginalVideoID(originalVideoID int) (*domain.FinalVideo, error)
}
