package main

// 导包方式
// import "fmt"
// import "time"
import (
	"fmt"
	"time"
)

func main() { // 注意：Go语言的大括号开头一定要在函数名后
	time.Sleep(1 * time.Second)
	fmt.Println("hello word")
}
