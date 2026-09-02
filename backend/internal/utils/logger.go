package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// InitLogger initializes Logrus structured logger
func InitLogger(env string) *logrus.Logger {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)

	if env == "production" {
		Log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		})
		Log.SetLevel(logrus.InfoLevel)
	} else {
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
		})
		Log.SetLevel(logrus.DebugLevel)
	}

	return Log
}

// GetLogger returns the global logger or a default one if not yet initialized
func GetLogger() *logrus.Logger {
	if Log == nil {
		return InitLogger("development")
	}
	return Log
}
