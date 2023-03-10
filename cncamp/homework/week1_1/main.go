package main

import "fmt"

//编写一个小程序：
//给定一个字符串数组
//[“I”,“am”,“stupid”,“and”,“weak”]
//用 for 循环遍历该数组并修改为
//[“I”,“am”,“smart”,“and”,“strong”]
func main() {
	arr := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Println("origin array is: ", arr)
	covert(&arr)
	fmt.Println("converted array is: ", arr)
}

// 修改原数组元素,入参需要传入指针类型
func covert(arr *[5]string) {
	for i := range arr {
		if arr[i] == "stupid" {
			arr[i] = "smart"
		}
		if arr[i] == "weak" {
			arr[i] = "strong"
		}
	}
}
