/*
channel的定义与使用
*/
package main

import "fmt"

func main() {
	// 创建一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine end...")
		fmt.Println("goroutine starting...")
		// 注意：当未被接收前会阻塞
		c <- 666
	}()

	// 注意：当为接收数据前会阻塞
	num := <-c
	fmt.Println("num = ", num)

	defer fmt.Println("main routine end...")
}
