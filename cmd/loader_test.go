package main

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	// Given
	testFileName := "test_config.toml"
	content := []byte("DBHost = \"localhost\"\n")
	err := os.WriteFile(testFileName, content, 0644)
	if err != nil {
		t.Fatalf("테스트 설정 파일을 생성할 수 없습니다: %v", err)
	}
	defer os.Remove(testFileName)

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{"cmd", "-c", testFileName}

	// When
	result := parseConfig()

	// Then
	assert.True(t, result, "parseConfig는 true를 반환해야 합니다.")
}

func TestLoadConfigWithNonExistingFile(t *testing.T) {
	// Given
	file = "non_existing_file.toml"

	// When
	result := loadConfig()

	// Then
	assert.False(t, result, "loadConfig는 false를 반환해야 합니다.")
}
