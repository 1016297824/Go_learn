/*
channel与range
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

	// for {
	// 	if num, ok := <-c; ok {
	// 		fmt.Println("get num ", num)
	// 	} else {
	// 		fmt.Println("channel close...")
	// 		break
	// 	}
	// }

	// 使用range简写
	for num := range c {
		fmt.Println("get num ", num)
	}

	fmt.Println("main end...")
}
