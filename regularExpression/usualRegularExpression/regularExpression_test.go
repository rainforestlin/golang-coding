package usualRegularExpression

import (
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
		{"AnyNumbers", "任意数字", "adadsjkl123", []string{}, []string{"123"}},
		{"nDigitalNumbers", "N位数字", "dfjk14234sklfjl12fjewe3456778", []string{"3", "5"}, []string{"14234", "34567"}},
		{"chinese", "中文", "my name is Julian Lee, chinese name 李俊霖", []string{}, []string{"李俊霖"}},
	}
	for _, test := range regularExpressionTest {
		actual := Verify(test.types, test.annotation, test.verifyString, test.ran)
		if !StringSliceEqual(actual, test.expected) {
			t.Errorf("Verfiy(%s) according to %s;get %s but expected %s", test.verifyString, test.annotation, actual, test.expected)
		}
	}
}

