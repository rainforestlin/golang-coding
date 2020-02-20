package main
//该例子用来学习接口的几个关键点
//1.接口是一种类型，从声明即可以看出使用type声明
//2.通过第四部分可以看出，接口变量存储的了实现者的值
//3.通过main函数可以看出，可以通过interface.(type)来判断接口变量存储的是哪个类型
//4.接口实现者的receiver的选择要根据实际情况，如果是值接收者，则传递值和指针都可以，如果是指针接收者则只能传递指针变量

import "fmt"


//第一部分  the first section
//声明一个接口
// Declare an interface
type Flyable interface {
	Fly() bool
	LearnToFly() bool
}
//第二部分 the second section
//声明一个struct
//Declare a struct
type Swift struct {
	Name      string
	IsFlyable bool
	CanSwim   bool
}

//声明另一个struct
//Declare another struct
type Dog struct {
	Name      string
	IsFlyable bool
	CanSwim   bool
}
//第三部分 the third section
// Swift实现了Flyable接口
// Swift implements Flyable
func (S Swift) Fly() bool {
	return S.IsFlyable
}

func (S Swift) LearnToFly() bool {
	S.IsFlyable = true
	return S.IsFlyable
}

// Dog实现了Flyable接口
//Dog implements Flyable
func (D Dog) Fly() bool {
	//狗不会飞
	return false
}

func (D *Dog) LearnToFly() bool {
	//狗学不会
	D.IsFlyable = false
	return D.IsFlyable
}
//第四部分 the fourth section
//定义两个函数
// Declare two functions to use interface
func VerifyIsFlyable(i Flyable) Flyable {
	switch i.(type) {
	case *Dog:
		fmt.Println("Dog can't fly forever")
	case *Swift:
		if i.(*Swift).IsFlyable {
			fmt.Println("Swift is the most fast flying speed bird in the world")
		} else {
			fmt.Println("May be she is too young")
		}
	case Swift:
		if i.(Swift).IsFlyable {
			fmt.Println("Swift is the most fast flying speed bird in the world")
		} else {
			fmt.Println("May be she is too young")
		}
	//	因为Dog实现的时候使用了指针接收者，所以无法使用值传递
	//case Dog:
	//	fmt.Println("Dog can't fly forever")
	}
	return i
}

func MakeItLearn(i Flyable) Flyable {
	i.LearnToFly()
	return i
}

func main() {
	var flyable Flyable
	newBornSwift := Swift{Name: "Minmin", IsFlyable: false, CanSwim: false}
	flyable = &newBornSwift
	if able, ok := VerifyIsFlyable(flyable).(*Swift); ok {
		fmt.Println("can her fly?", able.IsFlyable)
	} else {
		panic("interface assertion error")
	}

	grownUpSwift := MakeItLearn(&newBornSwift)
	if able, ok := VerifyIsFlyable(grownUpSwift).(*Swift); ok {
		fmt.Println("can her fly?", able.IsFlyable)
	} else {
		panic("interface assertion error")
	}

	newBornDog := Dog{Name: "Linlin", IsFlyable: false, CanSwim: true}
	flyable = &newBornDog
	if able, ok := VerifyIsFlyable(flyable).(*Dog); ok {
		fmt.Println("can him fly?", able.IsFlyable)
	} else {
		panic("interface assertion error")
	}
	grownUpDog := MakeItLearn(&newBornDog)
	if able, ok := VerifyIsFlyable(grownUpDog).(*Dog); ok {
		fmt.Println("can him fly?", able.IsFlyable)
	} else {
		panic("interface assertion error")
	}

	newBornSwift = Swift{Name: "Minmin2", IsFlyable: false, CanSwim: false}
	flyable = newBornSwift
	if able, ok := VerifyIsFlyable(flyable).(Swift); ok {
		fmt.Println("can her fly?", able.IsFlyable)
	} else {
		panic("interface assertion error")
	}
}
