package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	stringConversion()
}

//按照字节维度打印字符串
func printStrInBytes(str string) {
	fmt.Printf("the length of string is %d\n", len(str))
	fmt.Println("the bytes in string are: ")
	//按照字节维度遍历字符串,得到的每一个元素都是一个二进制字节
	for i := 0; i < len(str); i++ {
		fmt.Printf("0X%X\t", str[i]) //以二进制形式打印
	}
	fmt.Println()
}

//按照字符维度打印字符串
func printStringInRune(str string) {
	fmt.Printf("the rune count in string is %d\n", utf8.RuneCountInString(str))
	fmt.Println("the runes in string are: ")

	//按照字符维度(rune)遍历字符串,得到的每一个元素都是一个rune,对应的是一个Unicode 字符集表中的码点（Code Point）
	for _, r := range str {
		fmt.Printf("0X%X\t", r) //以二进制形式打印
	}
	fmt.Println()
}

//对Rune进行UTF-8编码
func encodeRune(r rune) {
	fmt.Printf("the rune is %c\n", r)
	buf := make([]byte, 3)
	utf8.EncodeRune(buf, r)
	fmt.Printf("the encode result is 0X%X\n", buf) //0XE4B8AD

}

//对Rune进行解码
func decodeRune(buf []byte) {
	fmt.Printf("the buf is 0X%X\n", buf)
	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("the decoded rune is %c\n", r)
}

//导出字符串类型变量的底层数组
func dumpUnderlyingArr(s string) {
	//将string类型变量的地址,显示转型为运行时结构reflect.StringHeader
	header := (*reflect.StringHeader)(unsafe.Pointer(&s))
	//获取底层数组
	fmt.Printf("the underlying array is 0X%X\t, len is %d\n", header.Data, header.Len)

	//获取底层数组的指针
	p := (*[5]byte)(unsafe.Pointer(header.Data))

	bytes := (*p)[:]

	fmt.Printf("[")
	for _, b := range bytes {
		fmt.Printf("%c\t", b)
	}
	fmt.Printf("]\n")
}

//字符串与字节、字符切片的转化
//直接显示类型转换即可
func stringConversion() {
	s := "中国人"

	//string -> []byte
	bs := []byte(s)
	fmt.Printf("string to byte slice is 0X%X\n", bs)

	//string -> []rune
	rs := []rune(s)
	fmt.Printf("string torune slice is 0X%X\n", rs)

	//[]byte -> string
	s1 := string(bs)
	fmt.Printf("byte slice to string is %s\n", s1)

	//[]rune -> string
	s2 := string(rs)
	fmt.Printf("rune slice to string is %s\n", s2)
}

//可以通过``定义原始字符串,打印结果就是所见即所得
func printRawStr() {
	var s string = `         ,_---~~~~~----._
    _,,_,*^____      _____*g*\"*,--,
   / __/ /'     ^.  /      \ ^@q   f
  [  @f | @))    |  | @))   l  0 _/
   \/   \~____ / __ \_____/     \
    |           _l__l_           I
    }          [______]           I
    ]            | | |            |
    ]             ~ ~             |
    |                            |
     |                           |`
	fmt.Println(s)
}
