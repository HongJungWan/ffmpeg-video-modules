package ffmpeg

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

type Job struct {
	Type       string
	InputPath  string
	OutputPath string
	StartTime  string
	EndTime    string
	InputPaths []string
}

func TrimVideo(inputPath string, outputPath string, startTime string, endTime string) error {
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

func ConcatVideos(inputPaths []string, outputPath string) error {
	fileList := "inputs.txt"
	f, err := os.Create(fileList)
	if err != nil {
		return fmt.Errorf("파일 목록 생성 실패: %w", err)
	}
	defer os.Remove(fileList)

	for _, path := range inputPaths {
		_, err := f.WriteString(fmt.Sprintf("file '%s'\n", path))
		if err != nil {
			return fmt.Errorf("파일 목록에 기록 실패: %w", err)
		}
	}
	f.Close()

	log.Printf("Executing ffmpeg command: Concat videos into %s", outputPath)
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
