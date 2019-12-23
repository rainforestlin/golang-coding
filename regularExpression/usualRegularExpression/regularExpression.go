package usualRegularExpression

import (
	"fmt"
	"regexp"
)

type RegularExpression struct {
	Expression string
	Type       string
	Annotation string
}

func (re RegularExpression) VerifyExpression(str string) interface{} {
	compile, e := regexp.Compile(re.Expression)
	if e != nil {
		panic(e)
	}
	if re.Type == "simpleInt" {
		match := compile.FindAllString(str, -1)
		fmt.Printf("method:%s -----%s", re.Annotation, match)
		fmt.Println()
		return match
	}
	return ""

}
