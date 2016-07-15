package main

import (
	"flag"
	"github.com/diegogub/esgo"
	"github.com/diegogub/esgo/store/arango"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

var r *esgo.CommandRouter

var (
	Arango = flag.String("arango", "", "aragno eventstore host")
	dev    = flag.Bool("dev", false, "dev mode")
)

type ExampleCmdHandler struct {
}

func (e ExampleCmdHandler) Commands() []string {
	return []string{
		"MakeExample",
	}
}

func (c ExampleCmdHandler) Deal(cmd *esgo.Command) (esgo.Eventer, *esgo.CommandResult) {

	var e esgo.Eventer
	var res *esgo.CommandResult

	switch cmd.Name {
	case "MakeExample":
		log.Println("ACA")
		event := &ExampleEventDone{}
		res = esgo.NewCommandResult(event)

		// set data into event
		cmd.SetEvent(event)

		// try to build event
		err := event.Build()

		// print it as example
		if err != nil {
			res.HasFailed(err)
		}

		e = event
		log.Println(">>>>", e)
	}

	return e, res
}

func HandleCommand(c *gin.Context) {
	//Get name from header
	cmdName := c.Request.Header.Get("Command")
	log.Println(cmdName)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	command := esgo.NewCommand("", cmdName, body)
	cres := r.Push(command)
	c.JSON(201, cres)
}

func main() {
	flag.Parse()
	es := arango.ArangoES{}
	arango.Dev = *dev
	arango.Init(*Arango)

	r = esgo.NewCommandRouter(es)

	//Add Example commands handler
	r.AddCommandHandler(ExampleCmdHandler{}, ExampleCmdHandler{}.Commands()...)

	r := gin.New()
	r.POST("/cmd/:group", HandleCommand)
	r.Run(":9090")
}
