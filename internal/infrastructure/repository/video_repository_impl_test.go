package repository_impl_test

import (
	"github.com/HongJungWan/ffmpeg-video-modules/internal/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/internal/infrastructure/repository"
	"github.com/HongJungWan/ffmpeg-video-modules/test/mocks"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestVideoRepository_Save(t *testing.T) {
	// Given
	db, mock := mocks.SetupMockDBForVideo(t)
	repo := repository_impl.NewVideoRepository(db)
	video := &domain.Video{
		Filename: "test_video.mp4",
		FilePath: "/videos/test_video.mp4",
		Duration: 120,
		Status:   domain.Uploaded,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `videos`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// When
	err := repo.Save(video)

	// Then
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVideoRepository_FindAll(t *testing.T) {
	// Given
	db, mock := mocks.SetupMockDBForVideo(t)
	repo := repository_impl.NewVideoRepository(db)

	rows := sqlmock.NewRows([]string{"id", "filename", "file_path", "duration", "status"}).
		AddRow(1, "test_video_1.mp4", "/videos/test_video_1.mp4", 120, domain.Uploaded).
		AddRow(2, "test_video_2.mp4", "/videos/test_video_2.mp4", 150, domain.Processed)
	mock.ExpectQuery("SELECT \\* FROM `videos`").
		WillReturnRows(rows)

	// When
	videos, err := repo.FindAll()

	// Then
	assert.NoError(t, err)
	assert.Len(t, videos, 2)
	assert.Equal(t, "test_video_1.mp4", videos[0].Filename)
	assert.Equal(t, "test_video_2.mp4", videos[1].Filename)
	assert.NoError(t, mock.ExpectationsWereMet())
}
