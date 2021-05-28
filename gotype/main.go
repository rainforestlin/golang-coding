package main

import (
	"fmt"
	"unsafe"
)

func getDataType(v interface{}) {
	fmt.Printf("该类型为%T,占用内存%d \n", v, unsafe.Sizeof(v))
}

func main() {
	var (
		boolType       bool
		intType        int
		int8Type       int8
		int16Type      int16
		int32Type      int32
		int64Type      int64
		uintType       uint
		uint8Type      uint8
		uint16Type     uint16
		uint32Type     uint32
		uint64Type     uint64
		uintptrType    uintptr
		float32Type    float32
		float64Type    float64
		complex64Type  complex64
		complex128Type complex128
		byteType       byte
		runeType       rune
		nullStructType struct{}
	)
	fmt.Printf("该类型为%T,占用内存%d \n", boolType, unsafe.Sizeof(boolType))
	fmt.Printf("该类型为%T,占用内存%d \n", intType, unsafe.Sizeof(intType))
	fmt.Printf("该类型为%T,占用内存%d \n", int8Type, unsafe.Sizeof(int8Type))
	fmt.Printf("该类型为%T,占用内存%d \n", int16Type, unsafe.Sizeof(int16Type))
	fmt.Printf("该类型为%T,占用内存%d \n", int32Type, unsafe.Sizeof(int32Type))
	fmt.Printf("该类型为%T,占用内存%d \n", int64Type, unsafe.Sizeof(int64Type))
	fmt.Printf("该类型为%T,占用内存%d \n", uintType, unsafe.Sizeof(uintType))
	fmt.Printf("该类型为%T,占用内存%d \n", uint8Type, unsafe.Sizeof(uint8Type))
	fmt.Printf("该类型为%T,占用内存%d \n", uint16Type, unsafe.Sizeof(uint16Type))
	fmt.Printf("该类型为%T,占用内存%d \n", uint32Type, unsafe.Sizeof(uint32Type))
	fmt.Printf("该类型为%T,占用内存%d \n", uint64Type, unsafe.Sizeof(uint64Type))
	fmt.Printf("该类型为%T,占用内存%d \n", uintptrType, unsafe.Sizeof(uintptrType))
	fmt.Printf("该类型为%T,占用内存%d \n", float32Type, unsafe.Sizeof(float32Type))
	fmt.Printf("该类型为%T,占用内存%d \n", float64Type, unsafe.Sizeof(float64Type))
	fmt.Printf("该类型为%T,占用内存%d \n", complex64Type, unsafe.Sizeof(complex64Type))
	fmt.Printf("该类型为%T,占用内存%d \n", complex128Type, unsafe.Sizeof(complex128Type))
	fmt.Printf("该类型为%T,占用内存%d \n", byteType, unsafe.Sizeof(byteType))
	fmt.Printf("该类型为%T,占用内存%d \n", runeType, unsafe.Sizeof(runeType))
	fmt.Printf("该类型为%T,占用内存%d \n", nullStructType, unsafe.Sizeof(nullStructType))

	memoryAlignment()

}

func memoryAlignment() {
	type T1 struct {
		a struct{}
		b int64
	}
	type T2 struct {
		b int64
		a struct{}
	}
	fmt.Printf("该类型为%T,占用内存%d ,偏移 %d\n", T1{}, unsafe.Sizeof(T1{}), unsafe.Alignof(T1{}))
	fmt.Printf("该类型为%T,占用内存%d ,偏移 %d\n", T2{}, unsafe.Sizeof(T2{}), unsafe.Alignof(T2{}))
}
