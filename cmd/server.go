package main

import (
	"github.com/HongJungWan/ffmpeg-video-modules/internal/helper"
	"github.com/HongJungWan/ffmpeg-video-modules/internal/infrastructure/router"
	"gorm.io/gorm"
	"net/http"
)

func startServer(db *gorm.DB) {
	routers := router.NewRouter(conf, db)

	server := &http.Server{
		Addr:    ":3031",
		Handler: routers,
	}
	err := server.ListenAndServe()
	if err != nil {
		helper.ErrorPanic(err)
	}
}
