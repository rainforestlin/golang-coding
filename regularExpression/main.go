package main

import (
	"fmt"
	"golang-coding/regularExpression/usualRegularExpression"
)

func main() {
	re:=usualRegularExpression.RegularExpression{"nDigitalNumbers","N",[]string{"3","5"}}
	verifyExpression, e := re.VerifyExpression("dfjk14234sklfjl12fjewe3456778")
	if e != nil{
		panic(e)
	}
	fmt.Print(verifyExpression)
}
