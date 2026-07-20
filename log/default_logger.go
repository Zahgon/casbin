package log

import (
	"io"
)

type DefaultLogger struct {
	output      io.Writer
	eventTypes  map[EventType]bool
	logCallback func(entry *LogEntry) error
}

func NewDefaultLogger() *DefaultLogger { _ = "STUB: not implemented"; return nil }

func (l *DefaultLogger) SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

func (l *DefaultLogger) SetEventTypes(eventTypes []EventType) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *DefaultLogger) OnBeforeEvent(entry *LogEntry) error { _ = "STUB: not implemented"; return nil }

func (l *DefaultLogger) OnAfterEvent(entry *LogEntry) error { _ = "STUB: not implemented"; return nil }

func (l *DefaultLogger) SetLogCallback(callback func(entry *LogEntry) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *DefaultLogger) writeLog(entry *LogEntry) error { _ = "STUB: not implemented"; return nil }
