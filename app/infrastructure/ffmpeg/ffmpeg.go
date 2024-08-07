package ffmpeg

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"os/exec"
)

type Job struct {
	Type       string
	InputPath  string
	OutputPath string
	StartTime  string
	EndTime    string
	InputPaths []string
}

func UploadVideo(filePath string, outputPath string) error {
	cmd := exec.Command("cp", filePath, outputPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("비디오 업로드 실패: %w", err)
	}
	return nil
}

func TrimVideo(inputPath string, outputPath string, startTime string, endTime string) error {
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

func ConcatVideos(inputPaths []string, outputPath string) error {
	// 입력 파일 경로를 저장할 임시 파일 생성
	fileList := "inputs.txt"
	f, err := os.Create(fileList)
	if err != nil {
		return fmt.Errorf("파일 목록 생성 실패: %w", err)
	}
	defer os.Remove(fileList) // 함수 종료 시 임시 파일 삭제

	// 각 입력 경로를 임시 파일에 기록
	for _, path := range inputPaths {
		_, err := f.WriteString(fmt.Sprintf("file '%s'\n", path))
		if err != nil {
			return fmt.Errorf("파일 목록에 기록 실패: %w", err)
		}
	}
	f.Close()

	// FFmpeg를 사용하여 파일 목록을 통해 비디오 연결
	err = ffmpeg.Input(fileList, ffmpeg.KwArgs{"f": "concat", "safe": "0"}).
		Output(outputPath, ffmpeg.KwArgs{"c": "copy"}).
		OverWriteOutput().
		Run()
	if err != nil {
		return fmt.Errorf("비디오 연결 실패: %w", err)
	}
	return nil
}

func DownloadVideo(videoPath string) ([]byte, error) {
	data, err := exec.Command("cat", videoPath).Output()
	if err != nil {
		return nil, fmt.Errorf("비디오 다운로드 실패: %w", err)
	}
	return data, nil
}

func GetVideoInfo(videoPath string) (string, error) {
	probe, err := ffmpeg.Probe(videoPath)
	if err != nil {
		return "", fmt.Errorf("비디오 정보 가져오기 실패: %w", err)
	}
	return probe, nil
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

func ConvertWebMToMP4(inputPath string, outputPath string) error {
	err := ffmpeg.Input(inputPath).
		Output(outputPath, ffmpeg.KwArgs{"c:v": "libx264", "c:a": "aac", "strict": "-2"}).
		OverWriteOutput().
		Run()
	if err != nil {
		return fmt.Errorf("WebM을 MP4로 변환하는 데 실패했습니다: %w", err)
	}
	return nil
}

func ConvertMP4ToHLS(inputPath string, outputDir string) error {
	err := ffmpeg.Input(inputPath).
		Output(fmt.Sprintf("%s/output.m3u8", outputDir), ffmpeg.KwArgs{
			"c:v":           "libx264",
			"c:a":           "aac",
			"hls_time":      10,
			"hls_list_size": 0,
			"f":             "hls",
		}).
		OverWriteOutput().
		Run()
	if err != nil {
		return fmt.Errorf("MP4를 HLS로 변환하는 데 실패했습니다: %w", err)
	}
	return nil
}
