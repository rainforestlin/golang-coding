package plugin

import (
	"fmt"
	"reflect"
)

var outputNames = make(map[string]reflect.Type)

type ConsoleOutput struct{}

func (c *ConsoleOutput) Send(str string) {
	fmt.Println(str)
}

func init() {
	outputNames["console"] = reflect.TypeOf(ConsoleOutput{})
}
