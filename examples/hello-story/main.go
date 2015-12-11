package main

import (
	"github.com/israelchen/gostory"
	"github.com/israelchen/gostory/handlers"
	"github.com/israelchen/gostory/rules"
)

func main() {

	fmtHandler := handlers.NewFmtHandler(gostory.DEBUG)
	story.AddHandler(rules.AlwaysOn, fmtHandler)

	s := story.New("blah")
	defer s.Done()

	s.Info("hello, story!").AddData("count", 123)
}
