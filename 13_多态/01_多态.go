/*
多态
*/
package main

import "fmt"

// 定义一个接口（本质是一个指针）
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 实现指针
type Cat struct {
	// 注意：此处不必调用接口，只要实现所有接口方法，即可指向接口
	color   string
	am_type string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping!!!")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return this.am_type
}

type Dog struct {
	// 注意：此处不必调用接口，只要实现所有接口方法，即可指向接口
	color   string
	am_type string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleeping!!!")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return this.am_type
}

func showAnimal(animalIF AnimalIF) {
	animalIF.Sleep()
	fmt.Println(animalIF.GetColor())
	fmt.Println(animalIF.GetType())
}

func main() {
	var animalIF AnimalIF
	animalIF = &Cat{"BaiSe", "Cat"}
	animalIF.Sleep()
	animalIF.GetColor()
	animalIF.GetType()

	dog := Dog{"HeiSe", "Dog"}
	showAnimal(&dog)
}
