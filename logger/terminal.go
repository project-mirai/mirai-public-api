package logger

import (
	"github.com/wsxiaoys/terminal"
	"github.com/wsxiaoys/terminal/color"
	"time"
)

type MiraiLogger struct {
	Prefix      string
	Terminal    *terminal.TerminalWriter
	ColorPrefix string
}

func NewLogger(prefix string) *MiraiLogger {
	return &MiraiLogger{
		Terminal: terminal.Stdout,
		Prefix:   prefix,
	}
}

func (this *MiraiLogger) Log(b string) {
	color.Print(this.ColorPrefix + "[" + this.Prefix + " | " + time.Now().Format("2006-01-02 15:04:05") + "]" + b + "\n")
}
