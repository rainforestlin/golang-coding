package testing

import (
	"golang-coding/regularExpression/usualRegularExpression"
	"testing"
)

func TestVerify(t *testing.T) {
	var regularExpressionTest = []struct {
		types        string
		annotation   string
		verifyString string
		ran          []string
		expected     []string
	}{
		{"AnyNumbers", "任意数字", "adadsjkl123", []string{},[]string{"123"}},
		//{"nDigitalNumbers", "N位数字", "dfjk14234sklfjl12fjewe3456778", []string{"3", "5"}},
	}
	for _, test := range regularExpressionTest {
		actual := usualRegularExpression.Verify(test.types, test.annotation, test.verifyString, test.ran)
		if StringSliceEqual(actual,test.expected){
			t.Errorf("Verfiy(%s) according to %s;get %s but expected %s",test.verifyString,test.annotation,actual,test.expected)

		}
	}
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}