package arrdemo

import (
	"fmt"
)

func ArrayTest() {
	// 全初始化，也可部分初始化var arr1 [3]int = [3]int{1, 3} //[1,3,0]
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)
	// 指定初始化[3,0,8]
	var arr2 [3]int = [3]int{0: 3, 2: 8}
	fmt.Println(arr2)

	// 定义的同时初始化,指定初始化,可以省略个数
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{3: 9}
	fmt.Print(arr3, arr4)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	// 在go语言中元素长度，类型，元素内容相同，两个数组才相同,且类型相同才能比较（数组长度也属于数组类型）
	// 如果元素类型支持==、!=操作时,那么数组也支持此操作
	// false
	fmt.Println(arr1 == arr2)
	// true
	fmt.Println(arr1 == arr3)
	// 编译报错
	// fmt.Println(arr1 == arr4)
}
