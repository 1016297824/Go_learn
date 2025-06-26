/*
结构体
*/
package main

import "fmt"

// type定义一种数据类型，myInt是int的别名
type myInt int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	book.title = "java"
}

func changeBook1(book *Book) {
	book.title = "java"
}

func main() {
	var a myInt = 2
	fmt.Printf("myInt is %T\n", a)

	// 结构体使用
	var book Book
	book.title = "Golang"
	book.auth = "ZK"
	fmt.Printf("book -> %v\n", book)

	// 结构体传参
	changeBook(book)
	fmt.Printf("book -> %v\n", book)
	changeBook1(&book)
	fmt.Printf("book -> %v\n", book)
}
