package plugin

type Plugin interface{}

type Input interface {
	Plugin
	Receive() string
}

type Filter interface {
	Plugin
	Process(str string) string
}

type Output interface {
	Plugin
	Send(str string)
}

type Type int

const (
	InputType = Type(iota)
	FilterType
	OutputType
)
