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
	docs.SwaggerInfo.Host = conf.Host
	docs.SwaggerInfo.Schemes = conf.Scheme
	docs.SwaggerInfo.Version = conf.Version
	docs.SwaggerInfo.BasePath = conf.BasePath
	docs.SwaggerInfo.Title = conf.Title

	fmt.Printf(
		"설정된 Swagger 정보:\nHost: %s\nSchemes: %v\nVersion: %s\nBasePath: %s\nTitle: %s\n",
		conf.Host,
		conf.Scheme,
		conf.Version,
		conf.BasePath,
		conf.Title,
	)
}
