package story

import (
	"fmt"
	"time"
	"util"
)

func (s *Story) Error(message string, args ...interface{}) *Story {
	return s.Log(ERROR, message, args...)
}

func (s *Story) Info(message string, args ...interface{}) *Story {
	return s.Log(INFO, message, args...)
}

func (s *Story) Warning(message string, args ...interface{}) *Story {
	return s.Log(WARNING, message, args...)
}

func (s *Story) Debug(message string, args ...interface{}) *Story {
	return s.Log(DEBUG, message, args...)
}

func (s *Story) Fatal(message string, args ...interface{}) *Story {
	return s.Log(FATAL, message, args...)
}

func (s *Story) Log(severity LogSeverity, message string, args ...interface{}) *Story {

	util.Require(len(message) > 0, "'message' cannot be empty.")

	formatted := fmt.Sprintf(message, args...)

	s.lock.Lock()
	defer s.lock.Unlock()

	s.LogEntries = append(s.LogEntries, LogEntry{
		Severity:   severity,
		CreateTime: time.Now(),
		Message:    formatted,
	})

	return s
}
