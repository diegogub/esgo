package esgo

import (
	"encoding/json"
	"errors"
	"github.com/satori/go.uuid"
	"time"
)

var (
	ConcurrencyError = errors.New("Concurrency error")
)

type CommandHandler interface {
	Deal(cmd *Command) (Eventer, CommandResult)
}

type Command struct {
	ID      string
	Version uint64
	Name    string
	Time    time.Time

	SessionID string
	UserID    string
	UserRole  string

	Test bool

	Data []byte
}

func NewCommand(id, name string, data []byte) *Command {
	var cmd Command
	if id == "" {
		cmd.ID = uuid.NewV4().String()
	}

	cmd.Time = time.Now().UTC()
	if len(data) == 0 {
		data = make([]byte, 0)
	}

	cmd.Data = data
	return &cmd
}

// every command result
type CommandResult struct {
	Err    error  `json:"-"`
	Error  bool   `json:"error"`
	ErrMsg string `json:"errorMsg"`

	Stream  string                 `json:"stream"`
	Version uint64                 `json:"version"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (c *Command) SetEvent(i interface{}) error {
	return json.Unmarshal(c.Data, &i)
}
