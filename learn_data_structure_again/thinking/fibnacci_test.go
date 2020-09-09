package thinking

import (
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{1, 0},
		{2, 1},
		{4, 2},
		{9, 21},
	}
	for _, test := range tests {
		result := fib(test.in)
		if result != test.out {
			t.Errorf("in :%d, should out:%d,actual out :%d", test.in, test.out, result)
		}
	}
}
