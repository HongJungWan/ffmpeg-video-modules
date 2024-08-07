package main

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/helper"
	"github.com/HongJungWan/ffmpeg-video-modules/app/infrastructure/configs"
	"github.com/HongJungWan/ffmpeg-video-modules/app/infrastructure/logger"
	"go.uber.org/zap"
	"os"
)

var (
	conf = configs.Config{}
	file string
)

func main() {
	if !parseConfig() {
		helper.ShowHelp()
		os.Exit(-1)
	}

	log := logger.ConfigureLogger(conf.Environment)
	logger.LogCurrentLevel(log)

	db := configs.ConnectionDB(&conf)
	log.Info("config", zap.Any("config", db))

	logger.LogCurrentConfig(conf)

	startServer()
}
