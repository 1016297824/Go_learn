/*
空接口
*/
package main

import "fmt"

// interface{}被称作具体类型，其它如int，float等被称作静态类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	// interface{}的类型断言机制
	// 注意：value存储的是断言类型
	value, ok := arg.(Book)
	if ok {
		fmt.Printf("arg is %T\n", value)
		fmt.Println("arg is Book, value = ", value)
	} else {
		fmt.Println("arg is not Book")
		// fmt.Printf("arg is %T\n", value)
		fmt.Printf("arg is %T\n", arg)
		fmt.Println("arg is ", arg)
	}
	fmt.Println()
}

type Book struct {
	title string
}

func main() {
	book := Book{"Golang"}
	// 不需要取地址
	myFunc(book)
	myFunc("Java")
	myFunc(666)
	myFunc(3.14159265)
}
