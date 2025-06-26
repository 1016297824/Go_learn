/*
	Go语言中的函数
*/
package main

import "fmt"

// 方式一：一般函数声明
func test1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 500
	return c
}

// 方式二：多返回值函数声明
func test2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 500
	return c, 666
}

// 方式三：多返回值函数声明，返回值有形参名称
// 注意：返回值默认为0
func test3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 直接给形参赋值
	r1 = 1000
	r2 = 2000
	return
}

// 方式四：多返回值函数声明，返回值有形参名称（变形：多返回值同类型缺省）
func test4(a string, b int) (r1, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 直接给形参赋值
	r1 = 3000
	r2 = 4000
	return
}

func main() {
	// 一般函数的使用
	result1 := test1("str", 100)
	fmt.Println("result = ", result1)
	println()

	// 多返回值函数的使用
	result2, result3 := test2("str2", 200)
	fmt.Println("result2 = ", result2)
	fmt.Println("result3 = ", result3)
	println()

	// 多返回值且有形参的函数使用
	result4, result5 := test3("str3", 300)
	fmt.Println("result4 = ", result4)
	fmt.Println("result5 = ", result5)
	println()

	// 多返回值且有形参的函数（变形）使用
	result6, result7 := test4("str4", 400)
	fmt.Println("result6 = ", result6)
	fmt.Println("result7 = ", result7)
	println()
}
