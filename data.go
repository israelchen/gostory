package gostory

import (
	"time"

	"github.com/israelchen/gostory/util"
)

func (s *Story) AddData(key string, value interface{}) *Story {

	util.Require(len(key) > 0, "'key' cannot be empty.")
	util.Require(s.HasEnded == false, "story is already done.")

	s.lock.Lock()
	defer s.lock.Unlock()

	s.Data[key] = DataEntry{
		CreateTime: time.Now(),
		Value:      value,
	}

	return s
}
