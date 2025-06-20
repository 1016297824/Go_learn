/*
变量声明方式
*/
package main

import "fmt"

func main() {
	// 方法一：声明一个变量（默认为：0）
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("Type of a = %T\n", a)

	// 方法二：声明并初始化一个变量
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("Type of b = %T\n", b)

	// 方法三：变量时赋值，自动匹配数据类型
	var c = "200"
	fmt.Println("c = ", c)
	fmt.Printf("Type of c = %T\n", c)

	// 方法四：省去var关键字，直接赋值并自动匹配（常用方法）
	// 注意：方法四不能声明全局变量
	d := 3.14159265
	fmt.Println("d = ", d)
	fmt.Printf("Type of d = %T\n", d)

	// 声明多个变量
	var zz, kk int = 300, 600
	fmt.Println("zz = ", zz, "kk = ", kk)
	var ww, yy = 300, "600"
	fmt.Println("ww = ", ww, "yy = ", yy)

	// 多行声明多个变量
	var (
		aa int  = 800
		bb bool = true
	)
	fmt.Println("aa = ", aa, "bb = ", bb)
}
