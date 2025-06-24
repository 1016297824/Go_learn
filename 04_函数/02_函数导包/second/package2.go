package package2

import "Go_learn/04_函数/02_函数导包/third"

func Package2Test() {
	println("package2Test()...")
}

func init() {
	println("package2 init()...")
	package3.Package3Test()
}
