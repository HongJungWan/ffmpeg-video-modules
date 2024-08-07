package main

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/helper"
	"github.com/HongJungWan/ffmpeg-video-modules/app/infrastructure/router"
	"net/http"
)

func startServer() {
	routers := router.NewRouter(conf)

	server := &http.Server{
		Addr:    ":3031",
		Handler: routers,
	}
	err := server.ListenAndServe()
	if err != nil {
		helper.ErrorPanic(err)
	}
}
