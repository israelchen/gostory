package handlers

import (
	"fmt"
	"story"
	"time"
)

func NewFmtHandler(severity story.LogSeverity) *fmtHandler {
	return &fmtHandler{
		severity: story.DEBUG,
	}
}

type fmtHandler struct {
	severity story.LogSeverity
}

func formatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

func (h *fmtHandler) Started(s *story.Story) {
}

func (h *fmtHandler) Stopped(s *story.Story) {

	shouldPrint := false

	for _, entry := range s.LogEntries {
		if entry.Severity >= h.severity {
			shouldPrint = true
			break
		}
	}

	if !shouldPrint {
		return
	}

	fmt.Printf("[%s][INFO] Story '%s' started.\n", formatTime(s.StartTime), s.Name)

	for _, entry := range s.LogEntries {
		fmt.Printf("[%s][%s] %s\n", formatTime(entry.CreateTime), story.LogSeverityNames[entry.Severity], entry.Message)
	}

	for key, value := range s.Data {
		fmt.Printf("[%s][DATA] %s => %v\n", formatTime(value.CreateTime), key, value.Value)
	}

	if s.Err != nil {
		fmt.Printf("[%s][ERROR] Story '%s' failed with error: %s.\n", formatTime(s.StartTime), s.Name, s.Err)
	} else {
		fmt.Printf("[%s][INFO] Story '%s' completed successfully.\n", formatTime(s.StartTime), s.Name)
	}
}
