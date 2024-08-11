package mocks

import (
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/stretchr/testify/mock"
)

type MockFinalVideoRepository struct {
	mock.Mock
}

func (m *MockFinalVideoRepository) FindFinalVideoByID(videoID int) (*domain.FinalVideo, error) {
	args := m.Called(videoID)
	return args.Get(0).(*domain.FinalVideo), args.Error(1)
}

func (m *MockFinalVideoRepository) SaveFinalVideo(finalVideo *domain.FinalVideo) error {
	args := m.Called(finalVideo)
	return args.Error(0)
}

func (m *MockFinalVideoRepository) FindFinalVideoByOriginalVideoID(originalVideoID int) (*domain.FinalVideo, error) {
	args := m.Called(originalVideoID)
	return args.Get(0).(*domain.FinalVideo), args.Error(1)
}

type MockVideoRepository struct {
	mock.Mock
}

func (m *MockVideoRepository) FindAll() ([]domain.Video, error) {
	args := m.Called()
	return args.Get(0).([]domain.Video), args.Error(1)
}

func (m *MockVideoRepository) Save(video *domain.Video) error {
	args := m.Called(video)
	return args.Error(0)
}

func (m *MockVideoRepository) FindByID(videoID int) (*domain.Video, error) {
	args := m.Called(videoID)
	return args.Get(0).(*domain.Video), args.Error(1)
}

type MockVideoJobRepository struct {
	mock.Mock
}

func (m *MockVideoJobRepository) Save(job *domain.VideoJob) error {
	args := m.Called(job)
	return args.Error(0)
}

func (m *MockVideoJobRepository) UpdateStatus(job *domain.VideoJob) error {
	args := m.Called(job)
	return args.Error(0)
}

func (m *MockVideoJobRepository) FindByID(videoID int) (*domain.VideoJob, error) {
	args := m.Called(videoID)
	return args.Get(0).(*domain.VideoJob), args.Error(1)
}

func (m *MockVideoJobRepository) FindByVideoIDAndType(videoID int, jobType domain.VideoJobType) ([]*domain.VideoJob, error) {
	args := m.Called(videoID, jobType)
	return args.Get(0).([]*domain.VideoJob), args.Error(1)
}

func (m *MockVideoJobRepository) FindPendingJobs() ([]*domain.VideoJob, error) {
	args := m.Called()
	return args.Get(0).([]*domain.VideoJob), args.Error(1)
}
