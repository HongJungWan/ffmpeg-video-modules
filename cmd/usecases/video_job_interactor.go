package usecases

import (
	"encoding/json"
	"fmt"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain/repository"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/infrastructure/ffmpeg"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/interfaces/dto/response"
	"os"
	"path/filepath"
)

type VideoJobInteractor interface {
	TrimVideo(videoID int, trimStart string, trimEnd string) (int, error)
	ConcatVideos(videoIDs []int) (int, error)
	ExecuteJobs(jobIDs []int) ([]response.JobIDResponse, error)
}

const (
	UPLOADS_DIR   = "uploads"
	DOWNLOADS_DIR = "downloads"
)

type VideoJobInteractorImpl struct {
	videoJobRepo   repository.VideoJobRepository
	videoRepo      repository.VideoRepository
	finalVideoRepo repository.FinalVideoRepository
}

func NewVideoJobInteractor(videoJobRepo repository.VideoJobRepository, videoRepo repository.VideoRepository, finalVideoRepo repository.FinalVideoRepository) VideoJobInteractor {
	return &VideoJobInteractorImpl{
		videoJobRepo:   videoJobRepo,
		videoRepo:      videoRepo,
		finalVideoRepo: finalVideoRepo,
	}
}

func (vji *VideoJobInteractorImpl) TrimVideo(videoID int, trimStart string, trimEnd string) (int, error) {
	video, err := vji.videoRepo.FindByID(videoID)
	if err != nil {
		return 0, fmt.Errorf("비디오를 찾을 수 없습니다: %w", err)
	}

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

func (vji *VideoJobInteractorImpl) ExecuteJobs(jobIDs []int) ([]response.JobIDResponse, error) {
	pendingJobs, err := vji.videoJobRepo.FindJobsByIDs(jobIDs)
	if err != nil {
		return nil, fmt.Errorf("작업을 불러올 수 없습니다: %w", err)
	}

	var successfulJobs []response.JobIDResponse
	var executionError error

	for _, job := range pendingJobs {
		job.UpdateStatus(domain.InProgress)
		if err := vji.videoJobRepo.UpdateStatus(job); err != nil {
			return nil, fmt.Errorf("작업 상태를 업데이트할 수 없습니다: %w", err)
		}

		var parameters map[string]interface{}
		if err := json.Unmarshal([]byte(job.Parameters), &parameters); err != nil {
			return nil, fmt.Errorf("작업 파라미터를 파싱할 수 없습니다: %w", err)
		}

		var execErr error
		var finalFilename string
		var finalDuration int

		switch job.JobType {
		case domain.Trim:
			inputPath := parameters["inputPath"].(string)
			outputPath := parameters["outputPath"].(string)
			trimStart := parameters["trimStart"].(string)
			trimEnd := parameters["trimEnd"].(string)

			execErr = ffmpeg.TrimVideo(inputPath, outputPath, trimStart, trimEnd)
			finalFilename = outputPath
			finalDuration, _ = ffmpeg.GetVideoDuration(outputPath)

		case domain.Concat:
			inputPathsInterface := parameters["inputPaths"].([]interface{})
			inputPaths := make([]string, len(inputPathsInterface))

			for i, path := range inputPathsInterface {
				inputPaths[i] = path.(string)
			}
			outputPath := parameters["outputPath"].(string)

			execErr = ffmpeg.ConcatVideos(inputPaths, outputPath)
			finalFilename = outputPath
			finalDuration, _ = ffmpeg.GetVideoDuration(outputPath)
		}

		if execErr != nil {
			job.UpdateStatus(domain.JobFailed)
			vji.videoJobRepo.UpdateStatus(job)
			originalVideo, _ := vji.videoRepo.FindByID(job.VideoID)
			originalVideo.UpdateStatus(domain.Failed)
			vji.videoRepo.Save(originalVideo)
			executionError = execErr
			continue
		} else {
			job.UpdateStatus(domain.Completed)
			successfulJobs = append(successfulJobs, response.JobIDResponse{JobID: job.ID})
		}

		originalVideo, _ := vji.videoRepo.FindByID(job.VideoID)
		finalVideo := domain.NewFinalVideo(originalVideo.ID, filepath.Base(finalFilename), finalFilename, finalDuration)
		if err := vji.finalVideoRepo.SaveFinalVideo(finalVideo); err != nil {
			return nil, fmt.Errorf("최종 비디오 정보를 저장할 수 없습니다: %w", err)
		}

		originalVideo.UpdateStatus(domain.FINISHED)
		vji.videoRepo.Save(originalVideo)

		if err := vji.videoJobRepo.UpdateStatus(job); err != nil {
			return nil, fmt.Errorf("작업 상태를 업데이트할 수 없습니다: %w", err)
		}
	}

	if executionError != nil {
		return successfulJobs, fmt.Errorf("일부 작업이 실패했습니다: %w", executionError)
	}

	return successfulJobs, nil
}
