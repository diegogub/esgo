package main

import (
	"github.com/diegogub/esgo"
	"github.com/diegogub/esgo/store/arango"
)

var r *esgo.CommandRouter

type ExampleCmdHandler struct {
}

func (e ExampleCmdHandler) Commands() []string {
	return []string{
		"MakeExaple",
	}
}

func (c ExampleCmdHandler) Deal(cmd *esgo.Command) (esgo.Eventer, *esgo.CommandResult) {

	var e esgo.Eventer
	var res *esgo.CommandResult

	switch cmd.Name {
	case "":
		event := &ExampleEventDone{}
		res = esgo.NewCommandResult(event)
		cmd.SetEvent(event)
		e = event
	}

	return e, res
}

func main() {
	es := arango.ArangoES{}
	arango.Init("")

	r = esgo.NewCommandRouter(es)

	//Add Example commands handler
	r.AddCommandHandler(ExampleCmdHandler{}, ExampleCmdHandler{}.Commands()...)

}
