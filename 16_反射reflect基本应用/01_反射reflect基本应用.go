/*
reflect反射机制
*/
package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("reflectNum is called...")
	fmt.Println("type = ", reflect.TypeOf(arg))
	fmt.Println("value = ", reflect.ValueOf(arg))
}

// 由于所有数据类型都是pari结构，所以，反射机制可以获取任意数据类型的值和类型
func main() {
	num := 3.14
	reflectNum(num)
}
