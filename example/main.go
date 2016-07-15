package main

import (
	"github.com/diegogub/esgo"
	"github.com/diegogub/esgo/store/arango"
)

var r *esgo.CommandRouter

type ExampleCmdHandler struct {
}

func (c ExampleCmdHandler) Deal(cmd *esgo.Command) (esgo.Eventer, esgo.CommandResult) {
}

func main() {
	es := arango.ArangoES{}
	arango.Init("")
	r = esgo.NewCommandRouter(es)

	r.AddCommandHandler()
}
