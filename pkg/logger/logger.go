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

func ErrorWithRequest(c *gin.Context, msg string, err error, req interface{}) {
	Logger.Error(msg,
		zap.Error(err),
		zap.Any("req_body", req),
		zap.String("client_ip", c.ClientIP()),
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
	)
}
