package usecases

import (
	"fmt"
	"github.com/HongJungWan/ffmpeg-video-modules/app/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/app/infrastructure/ffmpeg"
	"github.com/HongJungWan/ffmpeg-video-modules/app/interfaces/dto"
	"github.com/HongJungWan/ffmpeg-video-modules/app/interfaces/repository"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

type VideoInteractor interface {
	HandleVideoUpload(ctx *gin.Context) error
}

type VideoInteractorImpl struct {
	VideoRepository repository.VideoRepository
}

func NewVideoInteractor(videoRepository repository.VideoRepository) *VideoInteractorImpl {
	return &VideoInteractorImpl{VideoRepository: videoRepository}
}

func (vi *VideoInteractorImpl) HandleVideoUpload(ctx *gin.Context) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return fmt.Errorf("잘못된 폼 데이터: %v", err)
	}
	files := form.File["files"]

	var videoResponses []dto.VideoResponse

	for _, file := range files {
		// 파일 저장
		dst := filepath.Join("uploads", file.Filename)
		if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
			return fmt.Errorf("디렉토리를 생성할 수 없음: %v", err)
		}
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			return fmt.Errorf("파일을 저장할 수 없음: %v", err)
		}

		// 동영상 길이 계산
		duration, err := ffmpeg.GetVideoDuration(dst)
		if err != nil {
			return fmt.Errorf("동영상 길이를 계산할 수 없음: %v", err)
		}

		// 새로운 비디오 엔트리 생성
		videoResponse, err := vi.processVideo(file.Filename, dst, duration)
		if err != nil {
			return err
		}
		videoResponses = append(videoResponses, *videoResponse)
	}

	return nil
}

func (vi *VideoInteractorImpl) processVideo(filename string, filePath string, duration int) (*dto.VideoResponse, error) {
	video := domain.NewVideo(filename, filePath, duration)
	if err := vi.VideoRepository.Save(video); err != nil {
		return nil, fmt.Errorf("비디오를 생성할 수 없음: %v", err)
	}

	// 비디오 응답 반환
	videoResponse := &dto.VideoResponse{
		ID:        video.ID,
		Filename:  video.Filename,
		FilePath:  video.FilePath,
		Duration:  video.Duration,
		Status:    string(video.Status),
		CreatedAt: video.CreatedAt,
		UpdatedAt: video.UpdatedAt,
	}
	return videoResponse, nil
}
