package thinking

import "testing"

func TestCommon(t *testing.T) {
	tests := []struct {
		strA   string
		strB   string
		result string
	}{
		{
			"13452439",
			"123456",
			"345",
		},
	}

	for _, test := range tests {
		result := getCommon(test.strA, test.strB)
		if result != test.result {
			t.Errorf("result should be %s,but get %s", test.result, result)
		}
	}
}
