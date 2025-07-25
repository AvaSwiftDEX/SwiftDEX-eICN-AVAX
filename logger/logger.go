package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// InitLogger 初始化全局日志实例
func InitLogger(logFile string) {
	log = logrus.New()

	// 设置日志格式
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置输
	if logFile != "" {

		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			log.Error("Failed to open log file:", err)
		} else {
			log.SetOutput(file)
		}
	} else {
		log.SetOutput(os.Stdout)
	}

	// 设置日志级别
	log.SetLevel(logrus.InfoLevel)
}

// GetLogger 返回全局日志实例
func GetLogger() *logrus.Logger {
	return log
}

// NewComponent 为组件创建一个新的日志实例
func NewComponent(component string) *logrus.Entry {
	return log.WithField("component", component)
}

// SetOutput 设置日志输出目标
func SetOutput(output io.Writer) {
	log.SetOutput(output)
}

// SetLevel 设置日志级别
func SetLevel(level logrus.Level) {
	log.SetLevel(level)
}
