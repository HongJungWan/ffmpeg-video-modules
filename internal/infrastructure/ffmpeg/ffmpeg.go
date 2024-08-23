package ffmpeg

import (
	"fmt"
	"log"
	"os"
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
	inputPath := filepath.ToSlash(filepath.Join(UPLOADS_DIR, inputFilename))
	outputPath := filepath.ToSlash(outputFilename)
	if !strings.HasPrefix(outputFilename, DOWNLOADS_DIR) {
		outputPath = filepath.ToSlash(filepath.Join(DOWNLOADS_DIR, outputFilename))
	}

	log.Printf("Executing ffmpeg command: Trim video from %s to %s", startTime, endTime)
	cmd := exec.Command(FFMPEG_PATH, "-i", inputPath, "-ss", startTime, "-to", endTime, "-c:v", "libx264", "-c:a", "aac", "-strict", "-2", outputPath, "-y")
	log.Printf("Compiled command: %s", strings.Join(cmd.Args, " "))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("비디오 자르기 실패: %w", err)
	}
	return nil
}

func ConcatVideos(inputFilenames []string, outputFilename string) error {
	listFilePath := filepath.ToSlash(filepath.Join(DOWNLOADS_DIR, "concat_list.txt"))
	fileList := ""
	for _, inputFile := range inputFilenames {
		// 파일 경로를 리스트에 추가할 때 백슬래시 대신 슬래시를 사용
		absPath, err := filepath.Abs(inputFile)
		if err != nil {
			return fmt.Errorf("파일의 절대 경로를 찾을 수 없음: %w", err)
		}
		fileList += fmt.Sprintf("file '%s'\n", filepath.ToSlash(absPath))
	}

	// 파일 리스트를 임시 파일로 저장
	err := os.WriteFile(listFilePath, []byte(fileList), 0644)
	if err != nil {
		return fmt.Errorf("파일 목록을 저장할 수 없음: %w", err)
	}

	// 출력 경로 수정
	// outputFilename이 이미 DOWNLOADS_DIR에 포함된 경우 중복된 폴더 생성 방지
	if !strings.HasPrefix(outputFilename, DOWNLOADS_DIR) {
		outputFilename = filepath.Join(DOWNLOADS_DIR, outputFilename)
	}
	outputPath := filepath.ToSlash(outputFilename)

	// 출력 디렉토리 존재 여부를 확인하고, 없으면 생성
	outputDir := filepath.Dir(outputPath)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			return fmt.Errorf("출력 디렉토리를 생성할 수 없음: %w", err)
		}
	}

	// ffmpeg 명령어를 실행하여 비디오를 합침
	log.Printf("Executing ffmpeg command: Concat videos into %s", outputPath)
	cmd := exec.Command(FFMPEG_PATH, "-f", "concat", "-safe", "0", "-i", listFilePath, "-c:v", "libx264", "-c:a", "aac", "-strict", "-2", "-movflags", "+faststart", outputPath, "-y")
	log.Printf("Compiled command: %s", strings.Join(cmd.Args, " "))
	err = cmd.Run()
	if err != nil {
		// 명령어 실패 시 출력 확인
		errorOutput, _ := cmd.CombinedOutput()
		log.Printf("FFmpeg error output: %s", string(errorOutput))
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
