package log

type Logger interface {
	SetEventTypes([]EventType) error

	OnBeforeEvent(entry *LogEntry) error

	OnAfterEvent(entry *LogEntry) error

	SetLogCallback(func(entry *LogEntry) error) error
}
