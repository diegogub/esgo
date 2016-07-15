package esgo

import (
	"github.com/satori/go.uuid"
)

type Eventer interface {
	GetEventID() string
	GetStreamID() string
	GetType() string
	GetVersion() uint64

	GetStreamGroup() string
	GetUserID() string

	MustCreate() bool
	CheckUniqueValue() []string
}

type BaseEvent struct {
	eventid     string
	eventstream string
	eventType   string
	version     uint64
}

func (be *BaseEvent) GetEventID() string {
	return be.eventid
}

func (be *BaseEvent) SetEventID(id ...string) {
	if len(id) == 0 {
		be.eventid = uuid.NewV4().String()
	}

	return
}

func (be *BaseEvent) GetStreamID() string {
	return be.eventstream
}

func (be *BaseEvent) SetStream(id string) {
	be.eventstream = id
	return
}

func (be *BaseEvent) GetType() string {
	return be.eventType
}

func (be *BaseEvent) SetType(id string) {
	be.eventType = id
	return
}

func (be *BaseEvent) GetVersion() uint64 {
	return be.version
}

func (be *BaseEvent) SetVersion(v uint64) {
	be.version = v
}
