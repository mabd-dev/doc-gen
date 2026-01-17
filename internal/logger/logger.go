package logger

import (
	"fmt"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Gray   = "\033[90m"
)

type Logger struct {
	Quiet   bool
	Verbose bool
}

func (l Logger) LogError(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s[ERROR]%s %s\n", Red, Reset, msg)
}

func (l Logger) LogWarn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s[WARN]%s %s\n", Yellow, Reset, msg)
}

func (l Logger) LogInfo(format string, args ...any) {
	if l.Quiet {
		return
	}
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s[INFO]%s %s\n", Blue, Reset, msg)
}

func (l Logger) LogDebug(format string, args ...any) {
	if !l.Verbose {
		return
	}
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s[DEBUG]%s %s\n", Gray, Reset, msg)
}
