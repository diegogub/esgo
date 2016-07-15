package main

import (
	"errors"
	"github.com/diegogub/esgo"
	"time"
)

var (
	InvalidExampleID = errors.New("invalid id for stream")
)

const ExampleDoneEvent = "ExampleDone"

type ExampleEventDone struct {
	esgo.BaseEvent
	ID          string    `json:"id"`
	ExampleData string    `json:"example"`
	Date        time.Time `json:"date"`
}

func (te ExampleEventDone) GetStreamGroup() string {
	return "examples"
}

func (te ExampleEventDone) GetUserID() string {
	return "go"
}

func (te ExampleEventDone) MustCreate() bool {
	return false
}

func (te ExampleEventDone) CheckUniqueValue() []string {
	return []string{}
}

func (e *ExampleEventDone) Build() error {
	if e.ID == "" {
		return InvalidExampleID
	}
	e.SetStream(e.ID)
	e.SetType(ExampleDoneEvent)

	e.Date = time.Now().UTC()
	return nil
}
