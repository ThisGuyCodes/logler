package logler

import (
	"io"
	"log"
)

type Logger struct {
	Logger    *log.Logger
	levelChan chan uint8
	getChan   chan uint8
	level     uint8
}

func (l *Logger) Set(level int) {
	if l.levelChan == nil {
		l.levelChan = make(chan uint8)
		l.getChan = make(chan uint8)
		go l.levelGo()
	}
	l.levelChan <- uint8(level)
}

func (l *Logger) get() int {
	level := int(<-l.getChan)
	return level
}

func (l *Logger) levelGo() {
	for {
		select {
		case l.level = <-l.levelChan:
		case l.getChan <- l.level:
		}
	}
}

func New(out io.Writer, prefix string, flag int, level int) *Logger {
	logger := &Logger{}
	logger.Logger = log.New(out, prefix, flag)
	logger.Set(level)
	return logger
}
