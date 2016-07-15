package esgo

import (
	"encoding/json"
	"errors"
	"github.com/satori/go.uuid"
	"time"
)

var (
	ConcurrencyError   = errors.New("Concurrency error")
	InvalidCommandName = errors.New("Invalid command name")
)

type CommandHandler interface {
	Deal(cmd *Command) (Eventer, *CommandResult)
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

func (c *Command) Validate() error {
	if c.ID == "" {
		c.ID = uuid.NewV4().String()
	}

	if c.Name == "" {
		return InvalidCommandName
	}

	return nil
}

func (c *Command) SetEvent(i interface{}) error {
	return json.Unmarshal(c.Data, &i)
}

// every command result
type CommandResult struct {
	Err    error  `json:"-"`
	Error  bool   `json:"error"`
	ErrMsg string `json:"errorMsg"`

	Stream      string                 `json:"stream"`
	Version     uint64                 `json:"version"`
	Correlation uint64                 `json:"correlation"`
	Data        map[string]interface{} `json:"data,omitempty"`
}

func NewCommandResult(e Eventer) *CommandResult {
	var cmdRes CommandResult
	cmdRes.Stream = e.GetStreamID()
	cmdRes.Version = e.GetVersion()
	cmdRes.Data = make(map[string]interface{})
	return &cmdRes
}

func (cmr *CommandResult) HasFailed(err error) {
	if err != nil {
		cmr.Err = err
		cmr.Error = true
		cmr.ErrMsg = err.Error()
	}
}
