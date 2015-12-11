package main

import (
	"story"
	"story/handlers"
	"story/rules"
)

func main() {

	fmtHandler := handlers.NewFmtHandler(story.DEBUG)
	story.AddHandler(rules.AlwaysOn, fmtHandler)

	s := story.New("blah")
	defer s.Done()

	s.Info("hello, story!").AddData("count", 123)
}
