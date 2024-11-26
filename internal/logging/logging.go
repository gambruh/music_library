// package logging should implement the Logger interface for the service.
// However, for the moment it just wraps and returns logrus.Logger
package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger creates a new instance of logrus.Logger configured for JSON logging
func NewLogger(logFile string) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// If a log file is provided, write logs to the file
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}
		logger.SetOutput(file)
	}

	logger.SetLevel(logrus.DebugLevel)
	return logger, nil
}
