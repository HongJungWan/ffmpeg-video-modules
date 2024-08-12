package usecases

import (
	"fmt"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/domain/repository"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/infrastructure/ffmpeg"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/interfaces/dto/response"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"time"
)

type VideoInteractor interface {
	HandleVideoUpload(ctx *gin.Context) ([]response.VideoResponse, error)
	GetVideoDetails() ([]response.VideoDetailResponse, error)
}

type VideoInteractorImpl struct {
	VideoRepository      repository.VideoRepository
	VideoJobRepository   repository.VideoJobRepository
	FinalVideoRepository repository.FinalVideoRepository
}

func NewVideoInteractor(videoRepo repository.VideoRepository, videoJobRepo repository.VideoJobRepository, finalVideoRepo repository.FinalVideoRepository) *VideoInteractorImpl {
	return &VideoInteractorImpl{
		VideoRepository:      videoRepo,
		VideoJobRepository:   videoJobRepo,
		FinalVideoRepository: finalVideoRepo,
	}
}

func (vdi *VideoInteractorImpl) GetVideoDetails() ([]response.VideoDetailResponse, error) {
	videos, err := vdi.VideoRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("비디오를 불러올 수 없습니다: %w", err)
	}

	var videoDetails []response.VideoDetailResponse
	for _, video := range videos {
		videoDetail := response.VideoDetailResponse{
			ID:        video.ID,
			Filename:  video.Filename,
			FilePath:  video.FilePath,
			Duration:  video.Duration,
			Status:    string(video.Status),
			CreatedAt: video.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: video.UpdatedAt.Format("2006-01-02 15:04:05"),
		}

		// 트림 작업 가져오기
		trimJobs, _ := vdi.VideoJobRepository.FindByVideoIDAndType(video.ID, domain.Trim)
		for _, job := range trimJobs {
			videoDetail.TrimJobs = append(videoDetail.TrimJobs, response.VideoJobDetail{
				ID:         job.ID,
				JobType:    string(job.JobType),
				Parameters: job.Parameters,
				Status:     string(job.Status),
			})
		}

		// 이어붙이기 작업 가져오기
		concatJobs, _ := vdi.VideoJobRepository.FindByVideoIDAndType(video.ID, domain.Concat)
		for _, job := range concatJobs {
			videoDetail.ConcatJobs = append(videoDetail.ConcatJobs, response.VideoJobDetail{
				ID:         job.ID,
				JobType:    string(job.JobType),
				Parameters: job.Parameters,
				Status:     string(job.Status),
			})
		}

		// 최종 동영상 가져오기
		finalVideo, err := vdi.FinalVideoRepository.FindFinalVideoByOriginalVideoID(video.ID)
		if err == nil && finalVideo != nil {
			videoDetail.FinalVideoDetail = &response.FinalVideoDetail{
				ID:           finalVideo.ID,
				Filename:     finalVideo.Filename,
				FilePath:     finalVideo.FilePath,
				DownloadLink: fmt.Sprintf("http://localhost:3031/%s", finalVideo.FilePath),
			}
		}

		videoDetails = append(videoDetails, videoDetail)
	}

	return videoDetails, nil
}

func (vi *VideoInteractorImpl) HandleVideoUpload(ctx *gin.Context) ([]response.VideoResponse, error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, fmt.Errorf("잘못된 폼 데이터: %v", err)
	}
	files := form.File["files"]

	var videoResponses []response.VideoResponse

	for _, file := range files {
		// 고유한 파일 이름 생성
		uniqueFilename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		dst := filepath.Join("uploads", uniqueFilename)
		log.Printf("Saving file to: %s", dst)

		// 업로드된 파일을 저장
		if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
			return nil, fmt.Errorf("디렉토리를 생성할 수 없음: %v", err)
		}
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			return nil, fmt.Errorf("파일을 저장할 수 없음: %v", err)
		}

		// 동영상 길이 계산
		log.Printf("Calculating duration for: %s", dst)
		duration, err := ffmpeg.GetVideoDuration(dst)
		if err != nil {
			return nil, fmt.Errorf("동영상 길이를 계산할 수 없음: %v", err)
		}

		// 새로운 비디오 엔트리 생성
		videoResponse, err := vi.processVideo(uniqueFilename, dst, duration)
		if err != nil {
			return nil, err
		}
		videoResponses = append(videoResponses, *videoResponse)
	}

	// 모든 비디오 응답을 반환
	return videoResponses, nil
}

func (vi *VideoInteractorImpl) processVideo(filename string, filePath string, duration int) (*response.VideoResponse, error) {
	video := domain.NewVideo(filename, filePath, duration)
	if err := vi.VideoRepository.Save(video); err != nil {
		return nil, fmt.Errorf("비디오를 생성할 수 없음: %v", err)
	}

	// 비디오 응답 반환
	videoResponse := &response.VideoResponse{
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
