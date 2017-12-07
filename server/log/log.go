package log

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/ktr0731/cris/server/config"
)

type Logger struct {
	*log.Logger
	out io.Writer
}

var logger *Logger

func NewLogger(config *config.Config) (*Logger, error) {
	if logger != nil {
		return logger, nil
	}

	var out io.Writer
	switch config.Logger.Output {
	case "stdout":
		out = os.Stdout
	case "stderr":
		out = os.Stderr
	case "file":
		f, err := os.Create(string(time.Now().Unix()))
		if err != nil {
			return nil, err
		}
		out = f
	}
	return &Logger{
		log.New(out, config.Logger.Prefix, log.Lshortfile|log.Ltime),
		out,
	}, nil
}

func (l *Logger) Close() error {
	if f, ok := l.out.(io.Closer); ok {
		return f.Close()
	}
	return nil
}
