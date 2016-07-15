package arango

import (
	"encoding/json"
	"errors"
	"github.com/diegogub/esgo"
	"gopkg.in/h2non/gentleman.v1/plugins/body"
)

var (
	ESError = errors.New("Failed to store event")
)

type ArangoEvent struct {
	ID          string                 `json:"_key,omitempty"`
	UserID      string                 `json:"user,omitempty"`
	Group       string                 `json:"group,omitempty"`
	Stream      string                 `json:"-"`
	Type        string                 `json:"type"`
	Version     uint64                 `json:"version,omitempty"`
	Create      bool                   `json:"create,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
	Checks      []string               `json:"checks,omitempty"`
	Correlation uint64                 `json:"correlation,omitempty"`
}

type ArangoStoreResponse struct {
	ID          string `json:"_key"`
	Version     uint64 `json:"version"`
	Correlation uint64 `json:"correlation"`
	Error       bool   `json:"error"`
	ErrType     string `json:"errType,omitempty"`
}

func (ae *ArangoEvent) SetData(b []byte) error {
	return json.Unmarshal(b, &ae.Data)
}

func (aes ArangoES) Store(event esgo.Eventer) esgo.StoreResult {
	var res esgo.StoreResult

	res.Stream = event.GetStreamID()

	sEvent := &ArangoEvent{
		ID:      event.GetEventID(),
		UserID:  event.GetUserID(),
		Group:   event.GetStreamGroup(),
		Stream:  event.GetStreamID(),
		Type:    event.GetType(),
		Version: event.GetVersion(),
		Create:  event.MustCreate(),
		Checks:  event.CheckUniqueValue(),
	}

	b, err := json.Marshal(event)
	if err != nil {
		res.Error = err
		return res
	}

	err = sEvent.SetData(b)
	if err != nil {
		res.Error = err
		return res
	}

	req := g.Request().
		Method("POST").
		Path(db+"/stream/"+sEvent.Stream).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Use(body.JSON(sEvent))

	r, err := req.Send()
	if err != nil {
		res.Error = err
	}

	if r.StatusCode == 201 {
		ares := ArangoStoreResponse{}
		err = r.JSON(&ares)
		if err != nil {
			res.Error = err
		}

		if ares.Error {
			res.Error = errors.New(ares.ErrType)
		} else {
			res.Version = ares.Version
			res.Correlation = ares.Correlation
		}
	} else {
		res.Error = ESError
	}

	return res
}
