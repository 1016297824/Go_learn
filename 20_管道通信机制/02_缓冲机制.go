/*
缓冲机制
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main starting...")
	defer fmt.Println("main end...")

	// 创建有缓冲的channel
	// 注意：当缓冲慢时，会阻塞发送
	c := make(chan int, 5)
	fmt.Println("channel len is ", len(c), " cap is ", cap(c))

	// 创建一个子goroutine
	go func() {
		fmt.Println("goroutine starting...")
		defer fmt.Println("goroutine end...")
		for i := 0; i < 20; i++ {
			c <- i
			fmt.Println("channel len is ", len(c), " cap is ", cap(c), " send msg is ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	// 接收子goroutine的消息
	for i := 0; i < 20; i++ {
		num := <-c
		fmt.Println("channel len is ", len(c), " cap is ", cap(c), " get msg is ", num)
		time.Sleep(2 * time.Second)
	}
}
