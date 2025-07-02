/*
channel与select
*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("get msg ", i)
		}
		// close(c)
		quit <- true
	}()

Exit:
	for i := 0; ; {
		select {
		case c <- i:
			i++
		case b := <-quit:
			if b {
				// 注意：单纯使用break无法跳出循环
				break Exit // 使用标签
			}
		}
	}

	fmt.Println("main end...")
}
