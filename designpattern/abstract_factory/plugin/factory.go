package plugin

import "reflect"

type Config struct {
	Name       string
	PluginType Type
}

type Factory interface {
	Create(conf Config) Plugin
}

type InputFactory struct{}

func (i *InputFactory) Create(conf Config) Plugin {
	t := inputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type FilterFactory struct{}

func (f *FilterFactory) Create(conf Config) Plugin {
	t := filterNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type OutputFactory struct{}

func (o *OutputFactory) Create(conf Config) Plugin {
	t := outputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}
