package logger

import (
	"fmt"
	"io"

	"go.uber.org/zap"
)

// Log is a global instance of Logger.
var Log *Logger

// Logger is a simple wrapper around zap.SugaredLogger.
type Logger struct {
	*zap.SugaredLogger
}

// New creates a new Logger.
func New(debug bool) (*Logger, error) {
	if Log != nil {
		return Log, nil
	}

	var (
		logger *zap.Logger
		err    error
	)

	if debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return nil, fmt.Errorf("new zap logger: %w", err)
	}

	Log = &Logger{logger.Sugar()}

	return Log, nil
}

// Writer returns the logger's io.Writer.
func (l *Logger) Writer() io.Writer {
	return zap.NewStdLog(l.Desugar()).Writer()
}
