package repository_impl_test

import (
	repository_impl "github.com/HongJungWan/ffmpeg-video-modules/cmd/infrastructure/repository"
	"github.com/HongJungWan/ffmpeg-video-modules/test/mocks"
	"testing"

	"github.com/DATA-DOG/go-sqlmock" // sqlmock 패키지 임포트
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/stretchr/testify/assert"
)

func TestFinalVideoRepository_SaveFinalVideo(t *testing.T) {
	// Given
	db, mock := mocks.SetupMockDBForFinalVideo(t)
	repo := repository_impl.NewFinalVideoRepository(db)
	finalVideo := &domain.FinalVideo{
		OriginalVideoID: 1,
		Filename:        "final.mp4",
		FilePath:        "/videos/final.mp4",
		Duration:        120,
		Status:          domain.Processed,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `final_videos`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// When
	err := repo.SaveFinalVideo(finalVideo)

	// Then
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFinalVideoRepository_FindFinalVideoByID(t *testing.T) {
	// Given
	db, mock := mocks.SetupMockDBForFinalVideo(t)
	repo := repository_impl.NewFinalVideoRepository(db)
	videoID := 1

	rows := sqlmock.NewRows([]string{"id", "filename", "file_path"}).
		AddRow(1, "final.mp4", "/videos/final.mp4")
	mock.ExpectQuery("SELECT \\* FROM `final_videos` WHERE `final_videos`.`id` = \\? ORDER BY `final_videos`.`id` LIMIT \\?").
		WithArgs(videoID, 1). // 여기서 두 번째 인수 1은 LIMIT의 값을 처리합니다.
		WillReturnRows(rows)

	// When
	video, err := repo.FindFinalVideoByID(videoID)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, video)
	assert.Equal(t, "final.mp4", video.Filename)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFinalVideoRepository_FindFinalVideoByOriginalVideoID(t *testing.T) {
	// Given
	db, mock := mocks.SetupMockDBForFinalVideo(t)
	repo := repository_impl.NewFinalVideoRepository(db)
	originalVideoID := 1

	rows := sqlmock.NewRows([]string{"id", "original_video_id", "filename"}).
		AddRow(1, originalVideoID, "final.mp4")
	mock.ExpectQuery("SELECT \\* FROM `final_videos` WHERE original_video_id = \\? ORDER BY `final_videos`.`id` LIMIT \\?").
		WithArgs(originalVideoID, 1).
		WillReturnRows(rows)

	// When
	video, err := repo.FindFinalVideoByOriginalVideoID(originalVideoID)

	// Then
	assert.NoError(t, err)
	assert.NotNil(t, video)
	assert.Equal(t, originalVideoID, video.OriginalVideoID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
