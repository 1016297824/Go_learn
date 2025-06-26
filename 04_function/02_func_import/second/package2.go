package package2

import "Go_learn/04_function/02_func_import/third"

func Package2Test() {
	println("package2Test()...")
}

func init() {
	println("package2 init()...")
	package3.Package3Test()
}
