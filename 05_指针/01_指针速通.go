/*
指针速通
*/
package main

import "fmt"

// 值传递
func change_def(a int, b int) {
	temp := a
	a = b
	b = temp
	return
}

// 指针传递
func change_pointer(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
	return
}

// Golange支持二级指针但很少使用
func main() {
	a := 10
	b := 20
	change_def(a, b)
	fmt.Println("a = ", a, " b = ", b)
	change_pointer(&a, &b)
	fmt.Println("a = ", a, " b = ", b)
}
