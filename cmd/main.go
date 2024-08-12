package main

import (
	"fmt"
	"github.com/HongJungWan/ffmpeg-video-modules/cmd/docs"
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
	initializeSwaggerHost(&conf)
	db := configs.ConnectionDB(&conf)
	startServer(db)
}

func initializeSwaggerHost(conf *configs.Config) {
	var host string
	var scheme []string

	host = conf.Host
	scheme = conf.Scheme

	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Schemes = scheme
	fmt.Printf("설정된 Swagger Host: %s, Schemes: %v \n", host, scheme)
}
