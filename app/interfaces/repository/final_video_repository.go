package repository

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/domain"
	"gorm.io/gorm"
)

type FinalVideoRepository interface {
	SaveFinalVideo(finalVideo *domain.FinalVideo) error
	FindFinalVideoByID(id int) (*domain.FinalVideo, error)
	FindFinalVideoByOriginalVideoID(originalVideoID int) (*domain.FinalVideo, error)
}

type FinalVideoRepositoryImpl struct {
	DB *gorm.DB
}

func NewFinalVideoRepository(db *gorm.DB) *FinalVideoRepositoryImpl {
	return &FinalVideoRepositoryImpl{DB: db}
}

func (fv *FinalVideoRepositoryImpl) SaveFinalVideo(finalVideo *domain.FinalVideo) error {
	return fv.DB.Create(finalVideo).Error
}

func (fv *FinalVideoRepositoryImpl) FindFinalVideoByID(id int) (*domain.FinalVideo, error) {
	var finalVideo domain.FinalVideo
	if err := fv.DB.First(&finalVideo, id).Error; err != nil {
		return nil, err
	}
	return &finalVideo, nil
}

func (fv *FinalVideoRepositoryImpl) FindFinalVideoByOriginalVideoID(originalVideoID int) (*domain.FinalVideo, error) {
	var finalVideo domain.FinalVideo
	err := fv.DB.
		Where("original_video_id = ?", originalVideoID).
		First(&finalVideo).Error
	if err != nil {
		return nil, err
	}
	return &finalVideo, nil
}
