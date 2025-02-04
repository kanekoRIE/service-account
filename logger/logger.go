package logger

import (
	"os"

	"github.com/snowzach/rotatefilehook"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()

	// ensure permission and existence of log file
	logFile, err := os.OpenFile("/var/log/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatal(err)
	}
	defer logFile.Close()

	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Logger.SetOutput(os.Stdout)

	Logger.SetLevel(logrus.InfoLevel)

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "/var/log/app.log",
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Level:      logrus.InfoLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	})

	if err != nil {
		logrus.Fatal(err)
	}

	// add custom hook to write log to both stdout and file
	Logger.AddHook(rotateFileHook)
}

func LogInfo(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Info(message)
}

func LogWarning(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Warning(message)
}

func LogError(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Error(message)
}

func LogCritical(message string, fields logrus.Fields) {
	Logger.WithFields(fields).Fatal(message)
}
