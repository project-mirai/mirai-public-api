package logger

import (
	"github.com/wsxiaoys/terminal"
	"time"
)

type MiraiLogger struct {
	Prefix   string
	Terminal *terminal.TerminalWriter
}

func NewLogger(prefix string) *MiraiLogger {
	return &MiraiLogger{
		Terminal: terminal.Stdout,
		Prefix:   prefix,
	}
}

func (this *MiraiLogger) Log(b string) {
	this.Terminal.Write([]byte("[" + this.Prefix + " | " + time.Now().Format("2006-01-02 15:04:05") + "]" + b + "\n"))
}
