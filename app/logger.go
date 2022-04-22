package app

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger

	// VerboseLogging - verbose logging
	VerboseLogging bool

	// QuietLogging - no logging
	QuietLogging bool
)

// Log returns the logger instance
func Log() *logrus.Logger {
	if log == nil {
		log = logrus.New()
		log.SetLevel(logrus.InfoLevel)
		if VerboseLogging {
			log.SetLevel(logrus.DebugLevel)
		}

		if QuietLogging {
			log.SetLevel(logrus.ErrorLevel)
		}

		log.Out = os.Stdout
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "15:04:05",
			ForceColors:     true,
		})
	}

	return log
}
