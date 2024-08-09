package repository

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/domain"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video *domain.Video) error
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
