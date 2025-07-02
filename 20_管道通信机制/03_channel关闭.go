/*
channel关闭
*/
package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("send msg ", i)
		}
		close(c)
	}()

	for {
		if num, ok := <-c; ok {
			fmt.Println("get num ", num)
		} else {
			fmt.Println("channel close...")
			break
		}
	}

	fmt.Println("main end...")
}
