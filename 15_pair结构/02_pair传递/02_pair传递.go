/*
pair传递
*/
package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read Book ...")
}

func (this *Book) WriteBook() {
	fmt.Println("Write Book ...")
}

func main() {
	// b:<type:Book,value:Book{}>
	b := Book{}

	// r:<type:null,value:null>
	var r Reader
	// r:<type:Book,value:&Book{}>
	r = &b
	r.ReadBook()

	var w Writer
	// w:<type:Book,value:&Book{}>
	// 由于w和r的type相同，所以w和r可以互相断言、转化
	w = r.(Writer)
	w.WriteBook()
}
