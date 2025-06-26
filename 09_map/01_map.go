/*
map用法
*/
package main

import "fmt"

func valChange(cityMap map[string]string) {
	cityMap["India"] = "DeLi"
}

func main() {

	// 第一种声明方式
	// 声明一个map，键是string类型，值是int类型
	// 注意：map没有容量（cap）
	var myMap map[string]int
	if myMap == nil {
		fmt.Println("map is null")
	}

	// 为map开辟空间
	myMap = make(map[string]int, 2)

	// 为map添加值
	myMap["one"] = 1
	myMap["tow"] = 2
	myMap["three"] = 3
	fmt.Println(myMap)

	// 第二种声明方式
	myMap1 := make(map[string]string)

	// 为map添加值
	myMap1["one"] = "C"
	myMap1["tow"] = "C++"
	myMap1["three"] = "Java"
	fmt.Println(myMap1)

	// 第三种声明方式
	myMap2 := map[string]bool{
		"true":  true,
		"false": false,
	}
	fmt.Println(myMap2)
	println()

	cityMap := map[string]string{
		"China":   "BeiJing",
		"America": "NewYork",
		"India":   "DeLi",
		"Japan":   "Tokyo",
	}
	// 遍历
	for key, val := range cityMap {
		println(key, ":", val)
	}
	println()

	// 删除
	delete(cityMap, "Japan")
	// 遍历
	for key, val := range cityMap {
		println(key, ":", val)
	}
	println()

	// 修改
	cityMap["India"] = "New DeLi"
	// 遍历
	for key, val := range cityMap {
		println(key, ":", val)
	}
	println()

	// 传参
	valChange(cityMap)
	// 遍历
	for key, val := range cityMap {
		println(key, ":", val)
	}
	println()
}
