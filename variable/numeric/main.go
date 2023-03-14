package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//printIntegerNumberLiteral()
	//printFloatNumberLiteral()

	useTypeAlias()
}

// MyIntTypeDefine Type Define类型定义
//使用Type Define定义的新类型与底层类型是完全不同的两种类型,无法直接互相赋值,需要进行显式类型转换
type MyIntTypeDefine int

//使用类型定义
func useTypeDefine() {
	//新定义的类型与底层类型是完全不同的两种类型,无法直接互相赋值,需要进行显式类型转换
	a := int32(1)
	b := MyIntTypeDefine(a)
	fmt.Println("b=", b)
}

// MyIntTypeAlias Type Alias类型别名
//使用Type Alias定义的新类型与底层类型完全等价,可以直接互相替代和直接赋值
type MyIntTypeAlias = int32

func useTypeAlias() {
	//使用Type Alias定义的新类型与底层类型完全等价,可以直接互相替代和直接赋值
	a := int32(6)
	var b MyIntTypeAlias
	b = a
	fmt.Println("b=", b)
}

//查看数值类型变量长度
func sizeOf() {
	//在X86-64平台架构下,下面的size均为8
	a, b, p := int(6), uint(6), uintptr(0x12345678)
	fmt.Println("size of signed integer is: ", unsafe.Sizeof(a))   //8
	fmt.Println("size of unsigned integer is: ", unsafe.Sizeof(b)) //8
	fmt.Println("size of unsigned pointer is: ", unsafe.Sizeof(p)) //8
}

//打印整型字面值
func printIntegerNumberLiteral() {
	a := 1_2_3       //十进制
	b := 0b1001_0001 //二进制
	c := 0o700       //八进制
	d := 0x3C_4D_8F  //十进制
	fmt.Printf("十进制: %d\n", a)
	fmt.Printf("二进制: %b\n", b)
	fmt.Printf("八进制: %O\n", c)
	fmt.Printf("十进制: %X\n", d)
}

//打印浮点型字面值
func printFloatNumberLiteral() {
	pi := float64(3.1415926)
	fmt.Printf("普通浮点型表示: %f\n", pi)
	fmt.Printf("十进制科学计数法表示: %e\n", pi)
	fmt.Printf("十六进制科学计数法表示: %X\n", pi)
}
