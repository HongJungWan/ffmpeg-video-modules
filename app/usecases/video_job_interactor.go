package usecases

import (
	"encoding/json"
	"fmt"

	"github.com/HongJungWan/ffmpeg-video-modules/app/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/app/infrastructure/ffmpeg"
	"github.com/HongJungWan/ffmpeg-video-modules/app/interfaces/repository"
)

type VideoJobInteractor interface {
	TrimVideo(videoID int, trimStart string, trimEnd string) (int, error)
	ConcatVideos(videoIDs []int) (int, error)
	ExecuteJobs() error
}

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
	_, err := vji.videoRepo.FindByID(videoID)
	if err != nil {
		return 0, fmt.Errorf("비디오를 찾을 수 없습니다: %w", err)
	}

	job := domain.NewVideoJob(videoID, domain.Trim, map[string]interface{}{
		"trimStart": trimStart,
		"trimEnd":   trimEnd,
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

	// 첫 번째 비디오 ID를 videoID로 설정
	videoID := videoIDs[0]

	job := domain.NewVideoJob(videoID, domain.Concat, map[string]interface{}{
		"videoIDs": videoIDs,
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
			inputPath, ok := parameters["inputPath"].(string)
			if !ok {
				return fmt.Errorf("inputPath 파라미터가 없거나 잘못되었습니다")
			}
			outputPath, ok := parameters["outputPath"].(string)
			if !ok {
				return fmt.Errorf("outputPath 파라미터가 없거나 잘못되었습니다")
			}
			trimStart, ok := parameters["trimStart"].(string)
			if !ok {
				return fmt.Errorf("trimStart 파라미터가 없거나 잘못되었습니다")
			}
			trimEnd, ok := parameters["trimEnd"].(string)
			if !ok {
				return fmt.Errorf("trimEnd 파라미터가 없거나 잘못되었습니다")
			}
			execErr = ffmpeg.TrimVideo(inputPath, outputPath, trimStart, trimEnd)
		case domain.Concat:
			inputPaths, ok := parameters["inputPaths"].([]interface{})
			if !ok || len(inputPaths) == 0 {
				return fmt.Errorf("inputPaths 파라미터가 없거나 잘못되었습니다")
			}
			strInputPaths := make([]string, len(inputPaths))
			for i, path := range inputPaths {
				strPath, ok := path.(string)
				if !ok {
					return fmt.Errorf("inputPaths 내부에 잘못된 데이터가 포함되어 있습니다")
				}
				strInputPaths[i] = strPath
			}
			outputPath, ok := parameters["outputPath"].(string)
			if !ok {
				return fmt.Errorf("outputPath 파라미터가 없거나 잘못되었습니다")
			}
			execErr = ffmpeg.ConcatVideos(strInputPaths, outputPath)
		}

		if execErr != nil {
			job.UpdateStatus(domain.JobFailed)
		} else {
			job.UpdateStatus(domain.Completed)
		}

		if err := vji.videoJobRepo.UpdateStatus(job); err != nil {
			return fmt.Errorf("작업 상태를 업데이트할 수 없습니다: %w", err)
		}
	}

	return nil
}
