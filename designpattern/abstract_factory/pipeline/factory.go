package pipeline

import "github.com/julianlee107/learning/designpattern/abstract_factory/plugin"

var pluginFactory = make(map[plugin.Type]plugin.Factory)

func factoryOf(t plugin.Type) plugin.Factory {
	factory := pluginFactory[t]
	return factory
}

func Of(conf Config) *Pipeline {
	p := &Pipeline{}
	p.input = factoryOf(plugin.InputType).Create(conf.Input).(plugin.Input)
	p.output = factoryOf(plugin.OutputType).Create(conf.Output).(plugin.Output)
	p.filter = factoryOf(plugin.FilterType).Create(conf.Filter).(plugin.Filter)
	return p
}

func init() {
	pluginFactory[plugin.InputType] = &plugin.InputFactory{}
	pluginFactory[plugin.OutputType] = &plugin.OutputFactory{}
	pluginFactory[plugin.FilterType] = &plugin.FilterFactory{}

}
