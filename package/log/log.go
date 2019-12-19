package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Info(message ...string) {
	if gin.DebugMode == gin.Mode() {
		logInit("info")
		logger.Info(message)
	}
}

func Error(message string) {
	logInit("error")
	logger.Error(message)
}

func logInit(level string) {
	now := time.Now()

	dirPath := filepath.Join("./logs", now.Format("20060102"))
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {

	}

	logFilename := fmt.Sprintf("./logs/%s/%s.log", now.Format("20060102"), level)
	file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
}
