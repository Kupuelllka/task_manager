package logger

import (
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type logLevel uint

const (
	errorLevel    logLevel = 0
	infoLevel     logLevel = 1
	debugLevel    logLevel = 2
	debugRawLevel logLevel = 3
)

// Log contain logger info
type Log struct {
	log   *log.Logger
	level logLevel
}

func (l *Log) SetLevel(level string) {
	switch level {
	case "error":
		l.level = errorLevel
	case "info":
		l.level = infoLevel
	case "debug":
		l.level = debugLevel
	case "debug-raw":
		l.level = debugRawLevel
	default:
		l.level = errorLevel
	}
}
func (l *Log) getLevel(level string) logLevel {
	ret := errorLevel
	switch level {
	case "error":
		ret = errorLevel
	case "info":
		ret = infoLevel
	case "debug":
		ret = debugLevel
	case "debug-raw":
		ret = debugRawLevel
	}
	return ret
}

// InitLogger create new logger
func NewLogger(target string, level string, file string) (*Log, error) {
	ret := &Log{}
	ret.level = ret.getLevel(level)
	switch target {
	case "file":
		f, _ := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		ret.log = log.New(f, "", log.Ldate|log.Ltime)
	case "all":
		f, _ := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		m := io.MultiWriter(f, os.Stdout)
		ret.log = log.New(m, "", log.Ldate|log.Ltime)
	case "console":
		ret.log = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}
	return ret, nil
}

// Info вывод сообщений уровня info
func (l *Log) Infof(s string, a ...any) {
	if l.level >= infoLevel {
		l.log.Printf("[INFO]: "+s, a...)
	}
}

// Error вывод сообщений уровня error
func (l *Log) Errorf(s string, a ...any) {
	if l.level >= errorLevel {
		l.log.Printf("[ERROR]: "+s, a...)
	}
}

// Debug вывод сообщений уровня debug
func (l *Log) Debugf(s string, a ...any) {
	if l.level >= debugLevel {
		l.log.Printf("[DEBUG]: "+caller(2)+" - "+s, a...)
	}
}

func (l *Log) Fatalf(s string, a ...any) {
	l.log.Fatalf("[FATAL]: "+s, a...)
}

// Info вывод сообщений уровня info
func (l *Log) Info(s string) {
	if l.level >= infoLevel {
		l.log.Println("[INFO]: " + s)
	}
}

// Error вывод сообщений уровня error
func (l *Log) Error(s string) {
	if l.level >= errorLevel {
		l.log.Println("[ERROR]: " + s)
	}
}

// Debug вывод сообщений уровня debug
func (l *Log) Debug(s string) {
	if l.level >= debugLevel {
		l.log.Println("[DEBUG]: " + caller(2) + " - " + s)
	}
}

func (l *Log) Fatal(s string) {
	l.log.Fatalln("[FATAL]: " + s)
}

// DebugRaw вывод сообщений уровня debug raw
func (l *Log) DebugRaw(s string) {
	if l.level >= debugRawLevel {
		l.log.Println(s)
	}
}

// Version версии и номера сборки программы
func (l *Log) Version(s string) {
	l.log.Println(s)
}

// caller возвращает номер строки и распложение файла где был вызван один из методов логгера.
func caller(depth int) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(depth+1, pc)
	frame, _ := runtime.CallersFrames(pc[:n]).Next()
	idxFile := strings.LastIndexByte(frame.File, '/')
	idx := strings.LastIndexByte(frame.Function, '/')
	idxName := strings.IndexByte(frame.Function[idx+1:], '.') + idx + 1

	return frame.File[idxFile+1:] + ":[" + strconv.Itoa(frame.Line) + "] - " + frame.Function[idxName+1:] + "()"
}
