package ffmpeg

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	FFMPEG_PATH   = "C:\\ffmpeg\\bin\\ffmpeg.exe" // 경로 수정
	UPLOADS_DIR   = "uploads"
	DOWNLOADS_DIR = "downloads" // 변경된 결과물 저장 디렉토리
)

func TrimVideo(inputFilename string, outputFilename string, startTime string, endTime string) error {
	inputPath := filepath.Join(UPLOADS_DIR, inputFilename)
	outputPath := outputFilename
	if !strings.HasPrefix(outputFilename, DOWNLOADS_DIR) {
		outputPath = filepath.Join(DOWNLOADS_DIR, outputFilename)
	}

	log.Printf("Executing ffmpeg command: Trim video from %s to %s", startTime, endTime)
	cmd := exec.Command(FFMPEG_PATH, "-i", inputPath, "-ss", startTime, "-to", endTime, "-c", "copy", outputPath, "-y")
	log.Printf("Compiled command: %s", strings.Join(cmd.Args, " "))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("비디오 자르기 실패: %w", err)
	}
	return nil
}

func ConcatVideos(inputFilenames []string, outputFilename string) error {
	// 입력 파일 목록을 ffmpeg 명령어에 맞게 작성합니다.
	fileList := "concat:" + strings.Join(inputFilenames, "|")

	// 출력 경로 설정
	outputPath := outputFilename
	if !strings.HasPrefix(outputFilename, DOWNLOADS_DIR) {
		outputPath = filepath.Join(DOWNLOADS_DIR, outputFilename)
	}

	// ffmpeg 명령어를 실행합니다.
	log.Printf("Executing ffmpeg command: Concat videos into %s", outputPath)
	cmd := exec.Command(FFMPEG_PATH, "-i", fileList, "-c", "copy", outputPath, "-y")
	log.Printf("Compiled command: %s", strings.Join(cmd.Args, " "))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("비디오 연결 실패: %w", err)
	}

	return nil
}

func GetVideoDuration(filePath string) (int, error) {
	// ffprobe 명령어로 변경
	cmd := exec.Command("ffprobe", "-i", filePath, "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %s\n", string(output))
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
