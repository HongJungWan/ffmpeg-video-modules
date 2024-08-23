package mocks

import (
	domain2 "github.com/HongJungWan/ffmpeg-video-modules/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockFinalVideoRepository struct {
	mock.Mock
}

func (m *MockFinalVideoRepository) FindFinalVideoByID(videoID int) (*domain2.FinalVideo, error) {
	args := m.Called(videoID)
	return args.Get(0).(*domain2.FinalVideo), args.Error(1)
}

func (m *MockFinalVideoRepository) SaveFinalVideo(finalVideo *domain2.FinalVideo) error {
	args := m.Called(finalVideo)
	return args.Error(0)
}

func (m *MockFinalVideoRepository) FindFinalVideoByOriginalVideoID(originalVideoID int) (*domain2.FinalVideo, error) {
	args := m.Called(originalVideoID)
	return args.Get(0).(*domain2.FinalVideo), args.Error(1)
}

type MockVideoRepository struct {
	mock.Mock
}

func (m *MockVideoRepository) FindAll() ([]domain2.Video, error) {
	args := m.Called()
	return args.Get(0).([]domain2.Video), args.Error(1)
}

func (m *MockVideoRepository) Save(video *domain2.Video) error {
	args := m.Called(video)
	return args.Error(0)
}

func (m *MockVideoRepository) FindByID(videoID int) (*domain2.Video, error) {
	args := m.Called(videoID)
	return args.Get(0).(*domain2.Video), args.Error(1)
}

type MockVideoJobRepository struct {
	mock.Mock
}

func (m *MockVideoJobRepository) Save(job *domain2.VideoJob) error {
	args := m.Called(job)
	return args.Error(0)
}

func (m *MockVideoJobRepository) UpdateStatus(job *domain2.VideoJob) error {
	args := m.Called(job)
	return args.Error(0)
}

func (m *MockVideoJobRepository) FindByID(videoID int) (*domain2.VideoJob, error) {
	args := m.Called(videoID)
	return args.Get(0).(*domain2.VideoJob), args.Error(1)
}

func (m *MockVideoJobRepository) FindByVideoIDAndType(videoID int, jobType domain2.VideoJobType) ([]*domain2.VideoJob, error) {
	args := m.Called(videoID, jobType)
	return args.Get(0).([]*domain2.VideoJob), args.Error(1)
}

func (m *MockVideoJobRepository) FindPendingJobs() ([]*domain2.VideoJob, error) {
	args := m.Called()
	return args.Get(0).([]*domain2.VideoJob), args.Error(1)
}

func (m *MockVideoJobRepository) FindJobsByIDs(jobIDs []int) ([]*domain2.VideoJob, error) {
	args := m.Called(jobIDs)
	return args.Get(0).([]*domain2.VideoJob), args.Error(1)
}
