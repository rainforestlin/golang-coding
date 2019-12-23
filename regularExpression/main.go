package main

import (
	"fmt"
	"golang-coding/regularExpression/usualRegularExpression"
)
func main()  {
	anyNumber := usualRegularExpression.RegularExpression{"\\d+","simpleInt","任意数字"}
	result:=anyNumber.VerifyExpression("adkjflsfj456")
	fmt.Println(result)
	nNumber := usualRegularExpression.RegularExpression{"\\d{2}","simpleInt","N位数字"}
	result = nNumber.VerifyExpression("xdi234dll45")
	fmt.Print(result)
	
}
