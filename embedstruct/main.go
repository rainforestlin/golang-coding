package main

import "fmt"

//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

type UserP struct {
	Name    string
	Gender  string
	Address *Address
}

func main() {
	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City:     "威海",
		},
	}

	user2 := UserP{
		Name:   "冲冲冲",
		Gender: "女",
		Address: &Address{
			Province: "山东",
			City:     "威海",
		},
	}
	user3 := user1
	user4 := user2
	fmt.Printf("user1=%#v\n", user1) //user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
	fmt.Printf("user2=%#v\n", user2)
	fmt.Printf("user3=%#v\n", user3) //user3=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
	fmt.Printf("user4=%#v\n", user4)
	user1.Address.City = "青岛"
	user2.Address.City = "济南"
	fmt.Printf("user1=%#v city=%s\n ", user1, user1.Address.City)
	fmt.Printf("user2=%#v city=%s\n", user2, user2.Address.City)	
	fmt.Printf("user3=%#v city=%s \n ", user3, user3.Address.City)
	fmt.Printf("user4=%#v city=%s\n", user4, user4.Address.City)
}
