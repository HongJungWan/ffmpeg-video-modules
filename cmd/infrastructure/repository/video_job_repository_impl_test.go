package repository_impl_test

import (
	repository_impl "github.com/HongJungWan/ffmpeg-video-modules/cmd/infrastructure/repository"
	"github.com/HongJungWan/ffmpeg-video-modules/test/mocks"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/stretchr/testify/assert"
)

func TestVideoJobRepository_Save(t *testing.T) {
	// Given
	db, mock := mocks.SetupMockDBForVideoJob(t)
	repo := repository_impl.NewVideoJobRepository(db)
	videoJob := &domain.VideoJob{
		VideoID:    1,
		JobType:    domain.Trim,
		Parameters: `{"start_time":"00:00:10", "end_time":"00:00:20"}`,
		Status:     domain.Pending,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `video_jobs`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// When
	err := repo.Save(videoJob)

	// Then
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVideoJobRepository_FindPendingJobs(t *testing.T) {
	// Given
	db, mock := mocks.SetupMockDBForVideoJob(t)
	repo := repository_impl.NewVideoJobRepository(db)

	rows := sqlmock.NewRows([]string{"id", "video_id", "job_type", "parameters", "status"}).
		AddRow(1, 1, domain.Trim, `{"start_time":"00:00:10", "end_time":"00:00:20"}`, domain.Pending).
		AddRow(2, 2, domain.Concat, `{"files":["video1.mp4", "video2.mp4"]}`, domain.Pending)
	mock.ExpectQuery("SELECT \\* FROM `video_jobs` WHERE status = \\?").
		WithArgs(domain.Pending).
		WillReturnRows(rows)

	// When
	jobs, err := repo.FindPendingJobs()

	// Then
	assert.NoError(t, err)
	assert.Len(t, jobs, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVideoJobRepository_FindByVideoIDAndType(t *testing.T) {
	// Given
	db, mock := mocks.SetupMockDBForVideoJob(t)
	repo := repository_impl.NewVideoJobRepository(db)
	videoID := 1
	jobType := domain.Trim

	rows := sqlmock.NewRows([]string{"id", "video_id", "job_type", "parameters", "status"}).
		AddRow(1, videoID, jobType, `{"start_time":"00:00:10", "end_time":"00:00:20"}`, domain.Pending)
	mock.ExpectQuery("SELECT \\* FROM `video_jobs` WHERE video_id = \\? AND job_type = \\?").
		WithArgs(videoID, jobType).
		WillReturnRows(rows)

	// When
	jobs, err := repo.FindByVideoIDAndType(videoID, jobType)

	// Then
	assert.NoError(t, err)
	assert.Len(t, jobs, 1)
	assert.Equal(t, jobType, jobs[0].JobType)
	assert.NoError(t, mock.ExpectationsWereMet())
}
