/*
继承
*/
package main

import "fmt"

// 父类
type Hero struct {
	name string
	sex  string
}

func (this *Hero) eat() {
	fmt.Println("Hero.eat()...")
}

func (this *Hero) walk() {
	fmt.Println("Hero.walk()...")
}

// 子类
type SuperMan struct {
	Hero
	level int
}

func (this *SuperMan) walk() {
	fmt.Println("SuperMan.walk()...")
}

func (this *SuperMan) fly() {
	fmt.Println("SuperMan.fly()...")
}

func main() {
	hero := Hero{"people", "man"}
	hero.eat()
	hero.walk()

	superMan := SuperMan{Hero{"ChaoRen", "wemon"}, 999999}
	// 不重写->调用父类方法
	superMan.eat()
	// 重写->调用子类方法
	superMan.walk()
	superMan.fly()
}
