package casbin

import (
	"github.com/casbin/casbin/v3/log"
	"github.com/casbin/casbin/v3/model"
)

func (e *Enforcer) onLogBeforeEvent(eventType log.EventType) *log.LogEntry {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) onLogAfterEvent(logEntry *log.LogEntry) { _ = "STUB: not implemented"; return }

func (e *Enforcer) onLogAfterEventWithError(logEntry *log.LogEntry, err error) {
	_ = "STUB: not implemented"
	return
}

func countModelRules(m model.Model) int { _ = "STUB: not implemented"; return 0 }

func (e *Enforcer) onLogBeforeEventInLoadPolicy() *log.LogEntry {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) onLogAfterEventInLoadPolicy(logEntry *log.LogEntry, newModel model.Model) {
	_ = "STUB: not implemented"
	return
}

func (e *Enforcer) onLogBeforeEventInSavePolicy() *log.LogEntry {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) onLogAfterEventInSavePolicy(logEntry *log.LogEntry) {
	_ = "STUB: not implemented"
	return
}

func (e *Enforcer) createEnforceLogEntry(rvals []interface{}) *log.LogEntry {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) onLogBeforeEventInEnforce(rvals []interface{}) *log.LogEntry {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) onLogAfterEventInEnforce(logEntry *log.LogEntry, allowed bool) {
	_ = "STUB: not implemented"
	return
}

func (e *Enforcer) logPolicyOperation(eventType log.EventType, sec string, rule []string, operation func() (bool, error)) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
