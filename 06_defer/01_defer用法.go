/*
defer相关用法
*/
package main

// defer类似与析构函数，在函数结束时执行

func test1() {
	println("defer test1...")
}

func test2() {
	println("defer test2...")
}

func test3() {
	println("defer test3...")
}

func test4() {
	println("defer test4...")
}

func returnFucn() int {
	println("return test...")
	return 0
}

func deferAndReturn() int {
	defer test4()

	return returnFucn()
}

func main() {
	// 知识点一：defer执行顺序 -> 多个defer语句以栈的顺序执行
	defer test1()
	defer test2()
	defer test3()

	// 知识点二：defer和return执行顺序
	deferAndReturn()
}
