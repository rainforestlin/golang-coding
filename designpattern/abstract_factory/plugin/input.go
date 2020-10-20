package plugin

import "reflect"

var inputNames = make(map[string]reflect.Type)

type HelloInput struct{}

func (h *HelloInput) Receive() string {
	return "Hello world"
}

func init() {
	inputNames["hello"] = reflect.TypeOf(HelloInput{})
}
