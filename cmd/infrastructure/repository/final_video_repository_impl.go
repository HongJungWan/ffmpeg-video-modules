package repository_impl

import (
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain/repository"
	"gorm.io/gorm"
)

type FinalVideoRepositoryImpl struct {
	DB *gorm.DB
}

func NewFinalVideoRepository(db *gorm.DB) repository.FinalVideoRepository {
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
	err := fv.DB.Where("original_video_id = ?", originalVideoID).First(&finalVideo).Error
	if err != nil {
		return nil, err
	}
	return &finalVideo, nil
}
