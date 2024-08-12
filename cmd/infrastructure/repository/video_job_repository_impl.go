package repository_impl

import (
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain/repository"
	"gorm.io/gorm"
)

type VideoJobRepositoryImpl struct {
	DB *gorm.DB
}

func NewVideoJobRepository(db *gorm.DB) repository.VideoJobRepository {
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

func (repo *VideoJobRepositoryImpl) FindByVideoIDAndType(videoID int, jobType domain.VideoJobType) ([]*domain.VideoJob, error) {
	var jobs []*domain.VideoJob
	err := repo.DB.Where("video_id = ? AND job_type = ?", videoID, jobType).Find(&jobs).Error
	return jobs, err
}

func (repo *VideoJobRepositoryImpl) FindJobsByIDs(ids []int) ([]*domain.VideoJob, error) {
	var jobs []*domain.VideoJob
	err := repo.DB.Where("id IN ?", ids).Find(&jobs).Error
	return jobs, err
}
