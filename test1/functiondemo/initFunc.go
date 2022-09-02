package functiondemo

import (
	"fmt"
)

/**
golang里面有两个保留的函数：
init函数（能够应用于所有的package）
main函数（只能应用于package main）
这两个函数在定义时不能有任何的参数和返回值
go程序会自动调用init()和main()，所以你不能在任何地方调用这两个函数
package main必须包含一个main函数, 但是每个package中的init函数都是可选的
一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，我们都强烈建议用户在一个package中每个文件只写一个init函数
单个包中代码执行顺序如下
main包-->常量-->全局变量-->init函数-->main函数-->Exit
*/
func InitTest() {
	fmt.Println("initTest")
}

// init函数用于处理当前文件的初始化操作, 在使用某个文件时的一些准备工作应该放到这里
func init() {
	fmt.Println("init func")
}
