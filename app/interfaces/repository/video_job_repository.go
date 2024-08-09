package repository

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/domain"
	"gorm.io/gorm"
)

type VideoJobRepository interface {
	Save(job *domain.VideoJob) error
	FindByID(id int) (*domain.VideoJob, error)
	FindPendingJobs() ([]*domain.VideoJob, error)
	UpdateStatus(job *domain.VideoJob) error
}

type VideoJobRepositoryImpl struct {
	DB *gorm.DB
}

func NewVideoJobRepository(db *gorm.DB) *VideoJobRepositoryImpl {
	return &VideoJobRepositoryImpl{DB: db}
}

func (repo *VideoJobRepositoryImpl) Save(job *domain.VideoJob) error {
	return repo.DB.Create(job).Error
}

func (repo *VideoJobRepositoryImpl) FindByID(id int) (*domain.VideoJob, error) {
	var job domain.VideoJob
	err := repo.DB.First(&job, id).Error
	return &job, err
}

func (repo *VideoJobRepositoryImpl) FindPendingJobs() ([]*domain.VideoJob, error) {
	var jobs []*domain.VideoJob
	err := repo.DB.Where("status = ?", domain.Pending).Find(&jobs).Error
	return jobs, err
}

func (repo *VideoJobRepositoryImpl) UpdateStatus(job *domain.VideoJob) error {
	return repo.DB.Model(&job).Update("status", job.Status).Error
}
