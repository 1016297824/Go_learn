/*
结构体中的标签
*/
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `info:"姓名" json:"name"`
	Sex  string `info:"性别" json:"sex"`
}

func findTag(arg interface{}) {
	t := reflect.TypeOf(&arg).Elem()
	// 获取结构体的字段数
	for i := 0; i < t.NumField(); i++ {
		// 获取第i个标签的值
		infoTag := t.Field(i).Tag.Get("info")
		jsonTag := t.Field(i).Tag.Get("json")
		fmt.Println("infoTag = ", infoTag, ", jsonTag = ", jsonTag)
	}
}

func main() {
	var user User
	findTag(user)
}
