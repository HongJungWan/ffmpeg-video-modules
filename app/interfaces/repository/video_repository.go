package repository

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/domain"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video *domain.Video) error
	FindByID(videoID int) (*domain.Video, error)
}

type VideoRepositoryImpl struct {
	DB *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryImpl {
	return &VideoRepositoryImpl{DB: db}
}

func (vr *VideoRepositoryImpl) Save(video *domain.Video) error {
	return vr.DB.Create(video).Error
}

func (vr *VideoRepositoryImpl) FindByID(videoID int) (*domain.Video, error) {
	var video domain.Video
	err := vr.DB.First(&video, videoID).Error
	return &video, err
}
