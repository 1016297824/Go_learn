/*
匿名及别名导包方法
*/
package main

import (
	// 匿名导包，防止不使用报错
	_ "Go_learn/04_function/02_func_import/first"
	// 别名导包（不使用会报错）
	pg2 "Go_learn/04_function/02_func_import/second"
	// 将包嵌入当前文件
	. "Go_learn/04_function/02_func_import/third"
)

func main() {
	pg2.Package2Test()
	Package3Test()
}
