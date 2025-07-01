/*
pair结构-><type:**,value:**>
*/
package main

import "fmt"

func main() {
	var a string
	a = "abd"

	//pair结构
	var allType interface{}
	allType = a
	value, _ := allType.(string)
	fmt.Println("value = ", value)
}
