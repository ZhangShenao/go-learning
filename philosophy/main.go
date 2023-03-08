package main

import (
	"fmt"
	"time"
)

func main() {
	//explicitTypeConversion()
	concurrency()
}

// Go语言设计哲学——显式
func explicitTypeConversion() {
	var a int32 = 1
	var b int64 = 2
	var c int
	c = int(a) + int(b) //显式类型转换
	fmt.Printf("c=%d\n", c)
}

// Go语言设计哲学——并发
func concurrency() {
	go fmt.Println("This is a Goroutine") //开启一个Goroutine,执行并发任务
	time.Sleep(time.Second * 2)
}
