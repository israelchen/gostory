package main

import (
	"story"
	"story/handlers"
	"story/rules"
)

func main() {

	fmtHandler := &handlers.FmtHandler{story.DEBUG}
	story.AddHandler(rules.AlwaysOn, fmtHandler)

	s := story.New("blah")
	defer s.Done()

	s.LogInfo("hello, story!").AddData("count", 123)
}
