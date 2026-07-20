package log

import "time"

type EventType string

const (
	EventEnforce      EventType = "enforce"
	EventAddPolicy    EventType = "addPolicy"
	EventRemovePolicy EventType = "removePolicy"
	EventLoadPolicy   EventType = "loadPolicy"
	EventSavePolicy   EventType = "savePolicy"
)

type LogEntry struct {
	IsActive bool

	EventType EventType

	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration

	Subject string

	Object string

	Action string

	Domain string

	Allowed bool

	Rules [][]string

	RuleCount int

	Error error
}
