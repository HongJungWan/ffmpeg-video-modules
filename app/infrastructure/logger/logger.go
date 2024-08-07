package logger

import (
	"github.com/HongJungWan/ffmpeg-video-modules/app/infrastructure/configs"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ConfigureLogger(environment string) *zap.Logger {
	logLevel := getLogLevel(environment)
	encoderConfig := getEncoderConfig()
	consoleSyncer := getConsoleSyncer()

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), consoleSyncer, logLevel),
	)

	return zap.New(core, zap.AddCaller())
}

func getLogLevel(environment string) zap.AtomicLevel {
	if environment == "prd" {
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	}
	return zap.NewAtomicLevelAt(zap.InfoLevel)
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func getConsoleSyncer() zapcore.WriteSyncer {
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(os.Stderr),
	)
}

func LogCurrentLevel(currentLog *zap.Logger) {
	log.Printf("현재 설정된 log level: %v \n", currentLog.Level())
}

func LogCurrentConfig(config configs.Config) {
	log.Printf("현재 설정된 Config: %v \n", config)
}
