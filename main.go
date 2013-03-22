package logler

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	NONE     = 0
	CRITICAL = 1
	ERROR    = 2
	WARNING  = 3
	NORMAL   = 4
	VERBOSE  = 5
	DEBUG    = 6
)

type Logger struct {
	logger    *log.Logger
	levelChan chan<- uint8
	getChan   <-chan uint8
	level     uint8
}

func (l *Logger) SetLevel(level int) {
	if l.levelChan == nil {
		go l.levelGo()
		for l.levelChan == nil {

		}
	}
	l.levelChan <- uint8(level)
}

func (l *Logger) getLevel() int {
	level := int(<-l.getChan)
	return level
}

func (l *Logger) levelGo() {
	levelChan := make(chan uint8)
	getChan := make(chan uint8)
	l.levelChan = levelChan
	l.getChan = getChan
	for {
		select {
		case l.level = <-levelChan:
		case getChan <- l.level:
		}
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.getLevel() >= CRITICAL {
		l.logger.Fatal(v...)
	} else {
		os.Exit(1)
	}
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.getLevel() >= CRITICAL {
		l.logger.Fatalf(format, v...)
	} else {
		os.Exit(1)
	}
}

func (l *Logger) Fatalln(v ...interface{}) {
	if l.getLevel() >= CRITICAL {
		l.logger.Fatalln(v...)
	} else {
		os.Exit(1)
	}
}

func (l *Logger) Flags() int {
	return l.Flags()
}

func (l *Logger) Panic(level int, v ...interface{}) {
	if l.getLevel() >= level {
		l.logger.Panic(v...)
	} else {
		panic(fmt.Sprint(v...))
	}
}

func (l *Logger) Panicf(level int, format string, v ...interface{}) {
	if l.getLevel() >= level {
		l.logger.Panicf(format, v...)
	} else {
		panic(fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Panicln(level int, v ...interface{}) {
	if l.getLevel() >= level {
		l.logger.Panicln(v...)
	} else {
		panic(fmt.Sprintln(v...))
	}
}

func (l *Logger) Prefix() string {
	return l.logger.Prefix()
}

func (l *Logger) Print(level int, v ...interface{}) {
	if l.getLevel() >= level {
		l.logger.Print(v...)
	}
}

func (l *Logger) Printf(level int, format string, v ...interface{}) {
	if l.getLevel() >= level {
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Println(level int, v ...interface{}) {
	if l.getLevel() >= level {
		l.logger.Println(v...)
	}
}

func (l *Logger) SetFlags(flag int) {
	l.logger.SetFlags(flag)
}

func (l *Logger) SetPrefix(prefix string) {
	l.logger.SetPrefix(prefix)
}

func New(out io.Writer, prefix string, flag int, level int) *Logger {
	logger := &Logger{}
	logger.logger = log.New(out, prefix, flag)
	logger.SetLevel(level)
	return logger
}
