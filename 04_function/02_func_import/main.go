// init()函数
// 编译器解析过程：imoprt package -> const -> var -> init() -> main() 依次深度优先遍历解析
// 注意：同级包的编译器解析顺序可能不受import顺序影响，而受包在文件夹中的位置影响
package main

import "Go_learn/04_function/02_func_import/second"
import "Go_learn/04_function/02_func_import/first"

func main() {
	package1.Package1Test()
	package2.Package2Test()
}
