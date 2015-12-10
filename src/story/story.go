package story

import (
	"sync"
	"time"
	"util"
)

type LogSeverity int

const (
	DEBUG LogSeverity = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	LogSeverityNames map[LogSeverity]string = map[LogSeverity]string{
		DEBUG:   "DEBUG",
		INFO:    "INFO",
		WARNING: "WARNING",
		ERROR:   "ERROR",
		FATAL:   "FATAL",
	}
)

type LogEntry struct {
	Severity   LogSeverity
	CreateTime time.Time
	Message    string
}

type DataEntry struct {
	CreateTime time.Time
	Value      interface{}
}

type Story struct {
	Name       string
	Id         string
	StartTime  time.Time
	EndTime    time.Time
	LogEntries []LogEntry
	Data       map[string]DataEntry
	Err        error
	HasEnded   bool
	Parent     *Story
	Children   []*Story
	lock       sync.RWMutex
}

func New(name string) *Story {
	util.Require(len(name) > 0, "'name' cannot be empty!")

	result := &Story{
		Name:       name,
		Id:         "",
		StartTime:  time.Now(),
		LogEntries: []LogEntry{},
		Data:       make(map[string]DataEntry),
		HasEnded:   false,
	}

	for _, ph := range rules {
		if ph.Rule(result) {
			ph.Handler.Stopped(result)
		}
	}

	return result
}

func (s *Story) Done() *Story {
	util.Require(s.HasEnded == false, "story is already done.")

	s.HasEnded = true
	s.EndTime = time.Now()

	for _, ph := range rules {
		if ph.Rule(s) {
			ph.Handler.Stopped(s)
		}
	}

	return s
}

func (s *Story) SetSuccess() *Story {

	util.Require(s.HasEnded == false, "story is already done.")
	util.Require(s.Err == nil, "story already contains an error.")

	return s
}

func (s *Story) SetError(err error) *Story {

	util.Require(s.HasEnded == false, "story is already done.")
	util.Require(s.Err == nil, "story already contains an error.")

	s.Err = err
	return s
}
