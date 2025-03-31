package logger

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	var err error
	Logger, err = config.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	zap.ReplaceGlobals(Logger)
	zap.RedirectStdLog(Logger)
}

func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}

func Infof(format string, args ...any) {
	Logger.Sugar().Infof(format, args...)
}

func Errorf(format string, args ...any) {
	Logger.Sugar().Errorf(format, args...)
}

func Panicf(format string, args ...any) {
	Logger.Sugar().Panicf(format, args...)
}

func InfoWithRequest(ctx *gin.Context, msg string, req any) {
	Logger.Info(msg,
		zap.Any("req_body", req),
		zap.String("client_ip", ctx.ClientIP()),
		zap.String("method", ctx.Request.Method),
		zap.String("path", ctx.Request.URL.Path),
	)
}

func ErrorWithRequest(ctx *gin.Context, msg string, err error, req any) {
	Logger.Error(msg,
		zap.Error(err),
		zap.Any("req_body", req),
		zap.String("client_ip", ctx.ClientIP()),
		zap.String("method", ctx.Request.Method),
		zap.String("path", ctx.Request.URL.Path),
	)
}
