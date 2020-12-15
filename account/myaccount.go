package main

import "fmt"

func main() {
	key := ""
	loop := true
	balance := 0.0
	note := ""
	currentIncome := 0.0
	expend := 0.0
	for loop {
		fmt.Println("-----------家庭收支账本-----------")
		fmt.Println("            1、收支明细 ")
		fmt.Println("            2、登记收入 ")
		fmt.Println("            3、登记支出 ")
		fmt.Println("            4、退出 ")
		fmt.Println("请输入（1-4）")
		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Printf("当前余额%.2f \n4", balance)
		case "2":
			fmt.Print("请输入收入情况：")
			fmt.Scanln(&currentIncome)
			fmt.Println()
			fmt.Print("收入来源：")
			fmt.Scanln(&note)
			balance += currentIncome
			fmt.Printf("当前存款:%.2f \n", balance)

		case "3":
			fmt.Print("请输入支出")
			fmt.Scanln(&expend)
			fmt.Println()
			balance -= expend
			fmt.Print("支出原因:")
			fmt.Scanln(&note)
			fmt.Printf("当前存款:%.2f \n", balance)
		case "4":
			fmt.Println("退出")
			loop = false
		default:
			fmt.Println("没有输入")
		}
		if !loop {
			break
		}
	}
}
