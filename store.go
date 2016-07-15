package esgo

type EventStore interface {
	Store(event Eventer) StoreResult
}

type StoreResult struct {
	Error       error
	Stream      string
	Version     uint64
	Correlation uint64
}
