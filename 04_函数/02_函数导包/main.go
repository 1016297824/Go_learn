// init()函数
// 编译器解析过程：imoprt package -> const -> var -> init() -> main() 依次深度优先遍历解析
package main

import (
	"Go_learn/04_函数/02_函数导包/package1"
	"Go_learn/04_函数/02_函数导包/package2"
)

func main() {
	package1.Package1Test()
	package2.Package2Test()
}
