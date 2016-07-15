package arango

import (
	"gopkg.in/h2non/gentleman.v1"
	"log"
	"net/url"
)

var g *gentleman.Client
var db string

type ArangoES struct {
}

func Init(u string) {
	g = gentleman.New()
	nw, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	db = nw.Path

	log.Println("Init arango eventstore with url:", nw.Host, " and mount point ", nw.Path)
	g.URL(u)
}
