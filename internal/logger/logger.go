package logger

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var (
	currentLevel = LevelInfo
	timeFormat   = "15:04:05"
	verboseMode  = false
)

func SetLevel(l Level) { currentLevel = l }
func GetLevel() Level  { return currentLevel }
func SetVerbose(v bool) {
	verboseMode = v
}
func IsVerbose() bool { return verboseMode }

func ParseLevel(s string) Level {
	switch strings.ToLower(s) {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warn", "warning":
		return LevelWarn
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	default:
		return LevelInfo
	}
}

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "?????"
	}
}

var levelColors = map[Level]*color.Color{
	LevelDebug: color.New(color.FgCyan),
	LevelInfo:  color.New(color.FgWhite),
	LevelWarn:  color.New(color.FgYellow),
	LevelError: color.New(color.FgRed),
	LevelFatal: color.New(color.FgRed, color.Bold),
}

func enabled(l Level) bool { return l >= currentLevel }

func emit(lvl Level, tag, msg string) {
	w := os.Stdout
	if lvl >= LevelWarn {
		w = os.Stderr
	}
	c := levelColors[lvl]
	ts := time.Now().Format(timeFormat)
	if IsVerbose() {
		caller := callerFrame(3)
		msg = fmt.Sprintf("%s | caller=%s", msg, caller)
	}
	c.Fprintf(w, "%s  %-5s  %-12s  %s\n", ts, lvl, tag, msg)
}

func callerFrame(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

func Debug(tag, msg string) {
	if enabled(LevelDebug) {
		emit(LevelDebug, tag, msg)
	}
}
func Debugf(tag, format string, a ...any) {
	if enabled(LevelDebug) {
		emit(LevelDebug, tag, fmt.Sprintf(format, a...))
	}
}
func Info(tag, msg string) {
	if enabled(LevelInfo) {
		emit(LevelInfo, tag, msg)
	}
}
func Infof(tag, format string, a ...any) {
	if enabled(LevelInfo) {
		emit(LevelInfo, tag, fmt.Sprintf(format, a...))
	}
}
func Warn(tag, msg string) {
	if enabled(LevelWarn) {
		emit(LevelWarn, tag, msg)
	}
}
func Warnf(tag, format string, a ...any) {
	if enabled(LevelWarn) {
		emit(LevelWarn, tag, fmt.Sprintf(format, a...))
	}
}

func Error(tag, msg string) error {
	if enabled(LevelError) {
		emit(LevelError, tag, msg)
	}
	return errors.New(msg)
}
func Errorf(tag, format string, a ...any) error {
	safeFormat := strings.ReplaceAll(format, "%w", "%v")
	if enabled(LevelError) {
		emit(LevelError, tag, fmt.Sprintf(safeFormat, a...))
	}
	return fmt.Errorf(format, a...)
}
func Fatal(tag, msg string)               { emit(LevelFatal, tag, msg); os.Exit(1) }
func Fatalf(tag, format string, a ...any) { Fatal(tag, fmt.Sprintf(format, a...)) }
