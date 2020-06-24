package factorymethod

import "testing"

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestPlusOperator_Result(t *testing.T) {
	var (
		factory OperatorFactory
	)
	factory = &PlusOperatorFactor{}
	if compute(factory, 1, 2) != 3 {
		t.Fatal()
	}
}
