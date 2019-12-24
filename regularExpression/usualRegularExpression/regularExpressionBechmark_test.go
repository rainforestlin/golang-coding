package usualRegularExpression

import "testing"

func BenchmarkVerify(b *testing.B) {

	b.ResetTimer()
	for i:=0;i<b.N;i++{
		Verify("AnyNumbers", "任意数字", "adadsjkl123", []string{})
	}

	//for _, test := range regularExpressionTest {
	//	Verify(test.types, test.annotation, test.verifyString, test.ran)
	//}

}