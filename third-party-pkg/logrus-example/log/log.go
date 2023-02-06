package log

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Entry

type LogConfig struct {
	LogPath         string
	LogLevel        logrus.Level
	LogType         string
	LogMaxAge       time.Duration
	LogRotationTime time.Duration
}

var config = &LogConfig{
	LogPath:         "logs/app.log",
	LogLevel:        logrus.InfoLevel,
	LogType:         "file",
	LogMaxAge:       30 * 24 * time.Hour,
	LogRotationTime: 24 * time.Hour,
}

type LogWriter interface {
	Flush()
	io.Writer
}

type logWriter struct {
	*os.File
}

func (l *logWriter) Flush() {
	l.Sync()
}

func newLogWriter() LogWriter {
	file, err := os.OpenFile(config.LogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file = os.Stderr
	}
	return &logWriter{file}
}

// log file rotate
type fileRotateWriter struct {
	*rotatelogs.RotateLogs
}

func (frw *fileRotateWriter) Flush() {
	frw.Close()
}

func newFileRotateWriter() LogWriter {
	logf, err := rotatelogs.New(
		config.LogPath+".%Y%m%d",
		// rotatelogs.WithLinkName(config.LogPath),
		rotatelogs.WithMaxAge(config.LogMaxAge),
		rotatelogs.WithRotationTime(config.LogRotationTime),
	)
	if err != nil {
		return &logWriter{os.Stderr}
	}

	return &fileRotateWriter{logf}
}

func init() {
	Log = logrus.NewEntry(logrus.New())

	if config.LogType == "file" {
		// Log.Logger.SetOutput(newLogWriter())
		Log.Logger.SetOutput(newFileRotateWriter())
	} else {
		Log.Logger.SetOutput(os.Stdout)
	}

	Log.Logger.SetLevel(config.LogLevel)

	Log.Logger.SetFormatter(&logrus.JSONFormatter{})
	Log.Logger.SetReportCaller(true) // 打印行号
}
