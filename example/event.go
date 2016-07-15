package main

import (
	"github.com/diegogub/esgo"
	"time"
)

type ExampleEventDone struct {
	esgo.BaseEvent
	ExampleData string    `json:"example"`
	Date        time.Time `json:"date"`
}

func (te ExampleEventDone) GetStreamGroup() string {
	return "testing"
}

func (te ExampleEventDone) GetUserID() string {
	return "go"
}

func (te ExampleEventDone) MustCreate() bool {
	return false
}
