package arango

import (
	"github.com/diegogub/esgo"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

type TestEvent struct {
	esgo.BaseEvent
	TestData string    `json:"tdata"`
	Time     time.Time `json:"date"`
}

func NewTestEvent(data string) *TestEvent {
	var te TestEvent
	te.SetType("golangTest")
	te.SetStream("go-testing")
	te.TestData = data
	te.Time = time.Now().UTC()
	return &te
}

func (te TestEvent) GetStreamGroup() string {
	return "testing"
}

func (te TestEvent) GetUserID() string {
	return "go"
}

func (te TestEvent) MustCreate() bool {
	return false
}

func (te TestEvent) CheckUniqueValue() []string {
	return []string{}
}

func TestArangoES(t *testing.T) {
	Init("http://localhost:8529/_db/_system/s")

	te := NewTestEvent("diego")

	es := ArangoES{}

	res := es.Store(te)
	log.Println(res)
	assert.Nil(t, res.Error)
}

func BenchmarkArangoEs(b *testing.B) {
	Init("http://localhost:8529/_db/_system/s")
	es := ArangoES{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		te := NewTestEvent("diego")
		es.Store(te)
	}
}
