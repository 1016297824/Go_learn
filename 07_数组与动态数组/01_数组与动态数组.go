/*
数组与动态数组
*/
package main

import "fmt"

// 注意：固定长度的数组传参为值传递
func printArray(myArray [5]int) {
	for index, value := range myArray {
		println((index + 1), " is ", value)
	}
	myArray[1] = 500
}

// 注意：动态数组传参时传递的是地址（不是引用传递，但表现出引用传递的行为）
func printSlic(myArray []int) {
	// "_"表示匿名，不使用
	i := 0
	for _, value := range myArray {
		println((i + 1), " is ", value)
		i++
	}
	myArray[i-1] = 500
}

func main() {
	// 固定长度的数组
	var myArray [10]int
	// 赋值声明
	myArray2 := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(myArray); i++ {
		println((i + 1), " is ", myArray[i])
	}
	println()

	// 使用range遍历数组
	for index, value := range myArray2 {
		println((index + 1), " is ", value)
	}
	println()

	// 固定长度的数组传参类型必须完全对应
	printArray(myArray2)
	println()
	for index, value := range myArray2 {
		println((index + 1), " is ", value)
	}
	println()

	// 动态数组声明
	myArrayslice := []int{1, 2, 3, 4, 5} // 动态数组，切片 slice
	// 动态数组传参
	fmt.Printf("slice type is %T\n", myArrayslice)
	fmt.Println(myArrayslice)
	printSlic(myArrayslice)
	println()
	for index, value := range myArrayslice {
		println((index + 1), " is ", value)
	}
	println()
}
