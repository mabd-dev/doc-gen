package logger

import (
	"fmt"
	"os"
)

type Logger struct {
	Quiet   bool
	Verbose bool
}

func (l Logger) LogInfo(format string, args ...any) {
	if !l.Quiet {
		fmt.Fprintf(os.Stderr, format+"\n", args...)
	}
}

func (l Logger) LogDebug(format string, args ...any) {
	if l.Verbose {
		fmt.Fprintf(os.Stderr, "[Debug] "+format+"\n", args...)
	}
}

func (l Logger) LogWarn(format string, args ...any) {
	if l.Verbose {
		fmt.Fprintf(os.Stderr, "[Warn] "+format+"\n", args...)
	}
}

func (l Logger) LogError(format string, args ...any) {
	if l.Verbose {
		fmt.Fprintf(os.Stderr, "[Error] "+format+"\n", args...)
	}
}
