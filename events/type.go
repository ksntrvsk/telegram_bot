package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(event Event) error
}

type TypeEvent int

const (
	Unknown TypeEvent = iota
	Message
)

type Event struct {
	Type TypeEvent
	Text string
	Meta interface{}
}
