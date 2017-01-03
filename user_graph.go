package main

import (
	"fmt"
	"time"
)

const logFormat = "user_id: %d, following: %d, flollowers: %d"

type UserGraph struct {
	User_id   int64 ",omitempty"
	Following []int64
	Followers []int64
	Got_at    time.Time
}

func (self *UserGraph) ToLogFormat() string {
	return fmt.Sprintf(logFormat,
		self.User_id, len(self.Following), len(self.Followers))
}
