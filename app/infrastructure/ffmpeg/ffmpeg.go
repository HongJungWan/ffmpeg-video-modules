package ffmpeg

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Job struct {
	Type       string
	InputPath  string
	OutputPath string
	StartTime  string
	EndTime    string
	InputPaths []string
}

const (
	FFMPEG_PATH          = "C:\\ffmpeg\\bin\\ffprobe"
	UPLOADS_DIR          = "uploads"
	EXECUTE_JOB_DONE_DIR = "execute_job_done"
)

func TrimVideo(inputFilename string, outputFilename string, startTime string, endTime string) error {
	inputPath := filepath.Join(UPLOADS_DIR, inputFilename)
	outputPath := filepath.Join(EXECUTE_JOB_DONE_DIR, outputFilename)
	log.Printf("Executing ffmpeg command: Trim video from %s to %s", startTime, endTime)
	err := ffmpeg.Input(inputPath).
		Trim(ffmpeg.KwArgs{"start": startTime, "end": endTime}).
		Output(outputPath).
		OverWriteOutput().
		Run()
	if err != nil {
		return fmt.Errorf("비디오 자르기 실패: %w", err)
	}
	return nil
}

func ConcatVideos(inputFilenames []string, outputFilename string) error {
	fileList := "inputs.txt"
	f, err := os.Create(fileList)
	if err != nil {
		return fmt.Errorf("파일 목록 생성 실패: %w", err)
	}
	defer os.Remove(fileList)

	for _, filename := range inputFilenames {
		path := filepath.Join(UPLOADS_DIR, filename)
		_, err := f.WriteString(fmt.Sprintf("file '%s'\n", path))
		if err != nil {
			return fmt.Errorf("파일 목록에 기록 실패: %w", err)
		}
	}
	f.Close()

	outputPath := filepath.Join(EXECUTE_JOB_DONE_DIR, outputFilename)
	log.Printf("Executing ffmpeg command: Concat videos into %s", outputFilename)
	err = ffmpeg.Input(fileList, ffmpeg.KwArgs{"f": "concat", "safe": "0"}).
		Output(outputPath, ffmpeg.KwArgs{"c": "copy"}).
		OverWriteOutput().
		Run()
	if err != nil {
		return fmt.Errorf("비디오 연결 실패: %w", err)
	}
	return nil
}

func ExecuteJobs(jobs []Job) error {
	for _, job := range jobs {
		switch job.Type {
		case "trim":
			err := TrimVideo(job.InputPath, job.OutputPath, job.StartTime, job.EndTime)
			if err != nil {
				return fmt.Errorf("자르기 작업 실패: %w", err)
			}
		case "concat":
			err := ConcatVideos(job.InputPaths, job.OutputPath)
			if err != nil {
				return fmt.Errorf("연결 작업 실패: %w", err)
			}
		default:
			return fmt.Errorf("알 수 없는 작업 유형: %s", job.Type)
		}
	}
	return nil
}

func GetVideoDuration(filePath string) (int, error) {
	cmd := exec.Command(FFMPEG_PATH, "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", filePath)
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("비디오 정보를 불러오지 못했습니다: %w", err)
	}

	// 문자열의 불필요한 공백 및 줄 바꿈 문자 제거
	durationString := strings.TrimSpace(string(output))

	// 문자열을 float64로 변환
	duration, err := strconv.ParseFloat(durationString, 64)
	if err != nil {
		return 0, fmt.Errorf("동영상 길이를 변환할 수 없습니다: %w", err)
	}

	return int(duration), nil
}
