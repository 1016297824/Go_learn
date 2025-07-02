/*
退出一个goroutine
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 使用匿名函数，并开启一个goroutine
	go func() {
		defer fmt.Println("func A defer...")

		func() {
			defer fmt.Println("func B defer...")

			// 使用goexit()函数退出一个goroutine
			runtime.Goexit()

			fmt.Println("func B use...")
		}()

		fmt.Println("func A use...")
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}
