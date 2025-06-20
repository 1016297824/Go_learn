/*
常量学习
*/
package main

import "fmt"

// 使用const关键字定义枚举类型
const (
	// 可以在const()中定义一个iota，每行的iota会累加1，第一行默认为0
	BeiJing   = 10 * iota // iota = 0
	ShangHai              // iota = 1
	GuangDong             // iota = 2
	ShenZhen              // iota = 3
)

// 注意：iota只能使用在const关键字，且不受其它影响
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, j
)

func main() {
	// 常量声明
	const len = 100
	fmt.Println("len = ", len)
	fmt.Println("BeiJing = ", BeiJing)
	fmt.Println("ShangHai = ", ShangHai)
	fmt.Println("GuangDong = ", GuangDong)
	fmt.Println("ShenZhen = ", ShenZhen)

	fmt.Println()

	fmt.Println("a = ", a, " b = ", b)
	fmt.Println("c = ", c, " d = ", d)
	fmt.Println("e = ", e, " f = ", f)
	fmt.Println("g = ", g, " h = ", h)
	fmt.Println("i = ", i, " j = ", j)
}
