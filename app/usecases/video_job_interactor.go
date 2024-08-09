package usecases

import (
	"encoding/json"
	"fmt"
	"github.com/HongJungWan/ffmpeg-video-modules/app/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/app/infrastructure/ffmpeg"
	"github.com/HongJungWan/ffmpeg-video-modules/app/interfaces/repository"
	"os"
	"path/filepath"
)

type VideoJobInteractor interface {
	TrimVideo(videoID int, trimStart string, trimEnd string) (int, error)
	ConcatVideos(videoIDs []int) (int, error)
	ExecuteJobs() error
}

const (
	UPLOADS_DIR   = "uploads"
	DOWNLOADS_DIR = "downloads" // 변경된 결과물 저장 디렉토리
)

type VideoJobInteractorImpl struct {
	videoJobRepo repository.VideoJobRepository
	videoRepo    repository.VideoRepository
}

func NewVideoJobInteractor(videoJobRepo repository.VideoJobRepository, videoRepo repository.VideoRepository) VideoJobInteractor {
	return &VideoJobInteractorImpl{
		videoJobRepo: videoJobRepo,
		videoRepo:    videoRepo,
	}
}

func (vji *VideoJobInteractorImpl) TrimVideo(videoID int, trimStart string, trimEnd string) (int, error) {
	video, err := vji.videoRepo.FindByID(videoID)
	if err != nil {
		return 0, fmt.Errorf("비디오를 찾을 수 없습니다: %w", err)
	}

	// downloads 폴더 경로 생성
	if err := os.MkdirAll(DOWNLOADS_DIR, os.ModePerm); err != nil {
		return 0, fmt.Errorf("다운로드 디렉토리를 생성할 수 없습니다: %w", err)
	}

	outputFilename := fmt.Sprintf("trimmed_%s", video.Filename)
	outputPath := filepath.Join(DOWNLOADS_DIR, outputFilename)

	job := domain.NewVideoJob(videoID, domain.Trim, map[string]interface{}{
		"inputPath":  video.Filename,
		"outputPath": outputPath,
		"trimStart":  trimStart,
		"trimEnd":    trimEnd,
	})
	if err := vji.videoJobRepo.Save(job); err != nil {
		return 0, fmt.Errorf("작업을 저장할 수 없습니다: %w", err)
	}

	return job.ID, nil
}

func (vji *VideoJobInteractorImpl) ConcatVideos(videoIDs []int) (int, error) {
	if len(videoIDs) == 0 {
		return 0, fmt.Errorf("비디오 ID 목록이 비어 있습니다")
	}

	// downloads 폴더 경로 생성
	if err := os.MkdirAll(DOWNLOADS_DIR, os.ModePerm); err != nil {
		return 0, fmt.Errorf("다운로드 디렉토리를 생성할 수 없습니다: %w", err)
	}

	var inputFilenames []string
	for _, videoID := range videoIDs {
		video, err := vji.videoRepo.FindByID(videoID)
		if err != nil {
			return 0, fmt.Errorf("비디오를 찾을 수 없습니다: %w", err)
		}
		inputFilenames = append(inputFilenames, filepath.Join(UPLOADS_DIR, video.Filename))
	}

	outputFilename := "concatenated_video.mp4"
	outputPath := filepath.Join(DOWNLOADS_DIR, outputFilename)

	job := domain.NewVideoJob(videoIDs[0], domain.Concat, map[string]interface{}{
		"inputPaths": inputFilenames,
		"outputPath": outputPath,
	})
	if err := vji.videoJobRepo.Save(job); err != nil {
		return 0, fmt.Errorf("작업을 저장할 수 없습니다: %w", err)
	}

	return job.ID, nil
}

func (vji *VideoJobInteractorImpl) ExecuteJobs() error {
	pendingJobs, err := vji.videoJobRepo.FindPendingJobs()
	if err != nil {
		return fmt.Errorf("작업을 불러올 수 없습니다: %w", err)
	}

	for _, job := range pendingJobs {
		job.UpdateStatus(domain.InProgress)
		if err := vji.videoJobRepo.UpdateStatus(job); err != nil {
			return fmt.Errorf("작업 상태를 업데이트할 수 없습니다: %w", err)
		}

		var parameters map[string]interface{}
		if err := json.Unmarshal([]byte(job.Parameters), &parameters); err != nil {
			return fmt.Errorf("작업 파라미터를 파싱할 수 없습니다: %w", err)
		}

		var execErr error
		switch job.JobType {
		case domain.Trim:
			inputPath := parameters["inputPath"].(string)
			outputPath := parameters["outputPath"].(string)
			trimStart := parameters["trimStart"].(string)
			trimEnd := parameters["trimEnd"].(string)

			execErr = ffmpeg.TrimVideo(inputPath, outputPath, trimStart, trimEnd)

		case domain.Concat:
			inputPathsInterface := parameters["inputPaths"].([]interface{})
			inputPaths := make([]string, len(inputPathsInterface))

			for i, path := range inputPathsInterface {
				inputPaths[i] = path.(string)
			}
			outputPath := parameters["outputPath"].(string)

			execErr = ffmpeg.ConcatVideos(inputPaths, outputPath)
		}

		if execErr != nil {
			job.UpdateStatus(domain.JobFailed)
			vji.videoJobRepo.UpdateStatus(job)
			return execErr
		} else {
			job.UpdateStatus(domain.Completed)
		}

		if err := vji.videoJobRepo.UpdateStatus(job); err != nil {
			return fmt.Errorf("작업 상태를 업데이트할 수 없습니다: %w", err)
		}
	}

	return nil
}
