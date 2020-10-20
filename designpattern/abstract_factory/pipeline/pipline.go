package pipeline

import "github.com/julianlee107/learning/designpattern/abstract_factory/plugin"

type Pipeline struct {
	input  plugin.Input
	filter plugin.Filter
	output plugin.Output
}

type Config struct {
	Input  plugin.Config
	Filter plugin.Config
	Output plugin.Config
}

func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}

func DefaultConfig() Config {
	return Config{
		Input: plugin.Config{
			Name:       "hello",
			PluginType: plugin.InputType,
		},
		Filter: plugin.Config{
			Name:       "upper",
			PluginType: plugin.FilterType,
		},
		Output: plugin.Config{
			Name:       "console",
			PluginType: plugin.OutputType,
		},
	}
}
