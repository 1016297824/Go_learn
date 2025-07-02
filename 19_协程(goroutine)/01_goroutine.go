/*
协程
*/
package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newGoroutine() {
	for i := 1; i <= 20; i++ {
		fmt.Println("new goroutine -> ", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	// 通过go关键字创建一个go-routine
	go newGoroutine()

	for i := 1; i <= 10; i++ {
		fmt.Println("main goroutine -> ", i)
		time.Sleep(1 * time.Second)
	}

	// 主goroutine结束，则子goroutine也结束
	return
}
