package usecases

import (
	"fmt"
	"github.com/HongJungWan/ffmpeg-video-modules/internal/domain/repository"
	"github.com/HongJungWan/ffmpeg-video-modules/internal/interfaces/dto/response"
	"path/filepath"
)

type FinalVideoInteractor interface {
	GetFinalVideoDownloadLink(videoID int) (*response.VideoDownloadResponse, error)
}

type FinalVideoInteractorImpl struct {
	FinalVideoRepository repository.FinalVideoRepository
}

func NewFinalVideoInteractor(finalVideoRepository repository.FinalVideoRepository) *FinalVideoInteractorImpl {
	return &FinalVideoInteractorImpl{
		FinalVideoRepository: finalVideoRepository,
	}
}

func (fvi *FinalVideoInteractorImpl) GetFinalVideoDownloadLink(videoID int) (*response.VideoDownloadResponse, error) {
	finalVideo, err := fvi.FinalVideoRepository.FindFinalVideoByID(videoID)
	if err != nil {
		return nil, fmt.Errorf("최종 동영상을 찾을 수 없음: %w", err)
	}

	// 파일의 로컬 경로를 다운로드 링크로 변환
	downloadLink := fmt.Sprintf("http://localhost:3031/%s", filepath.ToSlash(finalVideo.FilePath))

	response := &response.VideoDownloadResponse{
		ID:           finalVideo.ID,
		Filename:     finalVideo.Filename,
		FilePath:     finalVideo.FilePath,
		DownloadLink: downloadLink,
	}

	return response, nil
}
