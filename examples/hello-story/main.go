package main

import (
	"github.com/israelchen/gostory"
	"github.com/israelchen/gostory/handlers"
	"github.com/israelchen/gostory/rules"
)

func main() {

	fmtHandler := handlers.NewFmtHandler(gostory.DEBUG)
	gostory.AddHandler(rules.AlwaysOn, fmtHandler)

	s := gostory.New("blah")
	defer s.Done()

	s.Info("hello, story!").AddData("count", 123)
}
