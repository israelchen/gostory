package story

import (
	"sync"
	"util"
)

type Handler interface {
	Started(s *Story)
	Stopped(s *Story)
}

type Rule func(*Story) bool

type PredicatedHandler struct {
	rule    Rule
	handler Handler
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
