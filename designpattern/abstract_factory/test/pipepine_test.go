package test

import (
	"github.com/julianlee107/learning/designpattern/abstract_factory/pipeline"
	"testing"
)

func TestPipeline(t *testing.T) {
	p := pipeline.Of(pipeline.DefaultConfig())
	p.Exec()
}
