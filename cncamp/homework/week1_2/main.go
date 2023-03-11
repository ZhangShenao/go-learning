package main

import (
	"fmt"
	"time"
)

//基于Channel编写一个简单的单线程生产者消费者模型：
//队列：
//队列长度10,队列元素类型为int
//生产者：
//每1秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
//消费者：
//每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
func main() {
	//define size of message queue
	queueSize := 10

	//create buffered channel as message queue
	queue := make(chan int, queueSize)

	//define done signal
	done := make(chan bool)

	//start producer
	producer(queue, done)

	//start consumer
	consumer(queue, done)

	//waiting for process done
	time.Sleep(time.Second * 10)

	//send done signal by closing channel
	close(done)
	time.Sleep(time.Second * 10)

	//exit main goroutine
	fmt.Println("[main]exit")
}

//Producer
func producer(queue <-chan int, done <-chan bool) {
	fmt.Println("producer started")
	go func() {
		//创建Ticker,每秒从队列中读取消息
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			select {
			case <-done:
				fmt.Println("[Consumer]process done")
				return
			case msg := <-queue:
				fmt.Printf("[Consumer]receive message: %d\n", msg)
			}
		}

	}()
}

//Consumer
func consumer(queue chan<- int, done <-chan bool) {
	fmt.Println("producer started...")
	go func() {
		//创建Ticker,每秒向队列中发送消息
		ticker := time.NewTicker(time.Second)
		counter := 1
		for range ticker.C {
			select {
			case <-done:
				fmt.Println("[Producer]process done")
				return
			case queue <- counter:
				fmt.Printf("[Producer]send message: %d\n", counter)
				counter += 1
			}
		}
	}()
}
