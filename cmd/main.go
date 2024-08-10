package main

import (
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/helper"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/infrastructure/configs"
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
	db := configs.ConnectionDB(&conf)
	startServer(db)
}
