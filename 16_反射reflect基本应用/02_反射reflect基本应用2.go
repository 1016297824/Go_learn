package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	println("User Call ...")
}

func (this *User) test() {
	println("User test ...")
}

func (this *User) End() {
	println("User End ...")
}

func DoFiledAndMethod(arg interface{}) {
	fmt.Println("DoFiledAndMethod ...")
	// 获取arg的type
	argType := reflect.TypeOf(arg)
	fmt.Println("argType = ", argType)
	// 获取arg的value
	argValue := reflect.ValueOf(arg)
	fmt.Println("argValue = ", argValue)

	// 循环获取argType中的字段和值
	// 每个字段中有属性：name和type
	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)
		// 注意：argValue中存储的属性不能是私有的！
		value := argValue.Field(i).Interface()
		fmt.Printf("field name = %s, field type = %v, field value = %v\n", field.Name, field.Type, value)
	}

	// 循环获取argType中的非指针方法
	for i := 0; i < argType.NumMethod(); i++ {
		method := argType.Method(i)
		fmt.Printf("method name = %s, method type = %v\n", method.Name, method.Type)
	}
}

func main() {
	u := User{1, "Tom", 18}
	u.Call()
	u.End()
	DoFiledAndMethod(u)
}
