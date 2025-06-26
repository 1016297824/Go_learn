/*
类
*/
package main

import "fmt"

// 定义类和方法
type Hero struct {
	name  string
	Ad    int
	Level int
}

func (this Hero) GetName() string {
	// fmt.Println("name is ", this.name)
	return this.name
}

func (this Hero) SetName1(name string) {
	// 注意：this是一个副本，不会修改原值
	this.name = name
	// fmt.Println("name is ", this.name)
}

func (this *Hero) SetName2(name string) {
	// 使用指针，修改原值
	this.name = name
	// fmt.Println("name is ", this.name)
}

func main() {
	hero := Hero{
		"ZZB",
		122,
		56,
	}
	fmt.Printf("Hero -> %v\n", hero)
	fmt.Println("name is ", hero.GetName())
	hero.SetName1("ZWZWZWZ")
	fmt.Println("name is ", hero.GetName())
	hero.SetName2("ZWZWZWZ")
	fmt.Println("name is ", hero.GetName())
}
