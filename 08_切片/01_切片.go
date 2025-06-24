/*
切片
*/
package main

import "fmt"

func main() {
	// 声明一个切片，并初始化
	mySlice1 := []int{1, 2, 3}
	fmt.Printf("slice length = %d\nslice = %v\n", len(mySlice1), mySlice1)
	println()

	// 声明一个切片，但不分配空间
	var mySlice2 []int
	fmt.Printf("slice length = %d\nslice = %v\n", len(mySlice2), mySlice2)
	if mySlice2 == nil {
		println("mySlice2 is null")
	}
	// 为mySlice2开辟三个空间
	mySlice2 = make([]int, 3)
	mySlice2[0] = 500
	fmt.Printf("slice length = %d\nslice = %v\n", len(mySlice2), mySlice2)
	if mySlice2 == nil {
		println("mySlice2 is null")
	} else {
		println("mySlice2 not null")
	}
	println()

	// 声明一个切片，同时分配空间
	var mySlice3 []int = make([]int, 3)
	fmt.Printf("slice length = %d\nslice = %v\n", len(mySlice3), mySlice3)
	println()

	// 直接使用make声明并分配空间
	mySlice4 := make([]int, 5)
	fmt.Printf("slice length = %d\nslice = %v\n", len(mySlice4), mySlice4)
	println()

	// 创建一个长度为3，容量为5的数组
	mySlice5 := make([]int, 3, 5)
	fmt.Printf("slice length = %d\nslice cap = %d\nslice = %v\n", len(mySlice5), cap(mySlice5), mySlice5)
	println()
	// 追加一个元素
	mySlice5 = append(mySlice5, 1)
	fmt.Printf("slice length = %d\nslice cap = %d\nslice = %v\n", len(mySlice5), cap(mySlice5), mySlice5)
	println()
	// 注意：追加元素超过当前容量时，会自动开辟初始容量的空间
	mySlice5 = append(mySlice5, 2)
	mySlice5 = append(mySlice5, 3)
	fmt.Printf("slice length = %d\nslice cap = %d\nslice = %v\n", len(mySlice5), cap(mySlice5), mySlice5)
	println()

	// 切片的截取
	// ???:容量截前不截后
	// 注意：此时截取的数组所有地址相同，一个改变，另一个也改变
	s1 := mySlice5[1:3]
	fmt.Printf("slice length = %d\nslice cap = %d\nslice = %v\n", len(s1), cap(s1), s1)
	println()

	// 切片的复制
	s2 := make([]int, 2)
	copy(s2, mySlice5)
	fmt.Printf("slice length = %d\nslice cap = %d\nslice = %v\n", len(s2), cap(s2), s2)
	println()
}
