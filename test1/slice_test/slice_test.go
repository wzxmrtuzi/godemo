package slicetest_test

import (
	"fmt"
)

// 无论是C语言中的数组还是Go语言中的数组,数组的长度一旦确定就不能改变, 但在实际开发中我们可能事先不能确定数组的长度
// 为了解决这类问题Go语言中推出了一种新的数据类型切片
// 切片可以简单的理解为长度可以变化的数组, 但是Go语言中的切片本质上是一个结构体
func SliceTest() {
	sliceCreate1()
}

// 通过数组创建切片
func sliceCreate1() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	// 从0下标开始，取值到下标2之前的下标，即，0，1
	var sce = arr1[0:2]
	// 切片len，结束位置-开始位置，2-0=2
	fmt.Println("len", len(sce))
	fmt.Println("cap", cap(sce))
}
