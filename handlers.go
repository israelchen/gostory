package gostory

import (
	"sync"

	"github.com/israelchen/gostory/util"
)

// It is strongly recommended to perform only quick
// and simple operations here and offload any long-running operation to
// a goroutine.
type Handler interface {
	// Started will be called as the story starts, in order to allow the handler
	// to perform some work.
	Started(s *Story)

	// Stopped will be called when a story ends, in order to allow the handler
	// to examine the different properties such as data, log, return error code.
	Stopped(s *Story)
}

// Rule provides a simple way to determine whether a handler applies
// to a given story.
type Rule func(*Story) bool

// PredicatedHandler couples a Handler with a Rule that determines whether
// it applies to a given story.
type PredicatedHandler struct {
	// Rule will determine whether the handler applies to
	// a given story.
	Rule Rule

	// A handler to invoke upon story start/stop.
	Handler Handler
}

var (
	hLock sync.RWMutex
	rules []PredicatedHandler = []PredicatedHandler{}
)

func AddHandler(rule Rule, handler Handler) {
	util.Require(handler != nil, "'handler' cannot be nil!")
	util.Require(rule != nil, "'rule' cannot be nil!")

	hLock.Lock()
	defer hLock.Unlock()
	rules = append(rules, PredicatedHandler{rule, handler})
}

func AddHandlers(handlers []PredicatedHandler) {
	util.Require(handlers != nil, "'handlers' cannot be nil!")

	hLock.Lock()
	defer hLock.Unlock()

	rules = append(rules, handlers...)
}

func RemoveAllHandlers() {
	hLock.Lock()
	defer hLock.Unlock()
	rules = []PredicatedHandler{}
}
