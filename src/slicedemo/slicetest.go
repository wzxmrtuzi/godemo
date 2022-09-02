package slicedemo

import (
	"fmt"
)

// 无论是C语言中的数组还是Go语言中的数组,数组的长度一旦确定就不能改变, 但在实际开发中我们可能事先不能确定数组的长度
// 为了解决这类问题Go语言中推出了一种新的数据类型切片
// 切片可以简单的理解为长度可以变化的数组, 但是Go语言中的切片本质上是一个结构体
func SliceTest() {
	sliceCreateByAssign()
	sliceCreateByMake()
	sliceCreateBySugar()
	sliceCopy()
}

// 通过数组创建切片
func sliceCreateByAssign() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	// 开始位置和结束位置都指定[1:2]
	// 只指定开始位置或结束位置[1:] [:2]
	// 开始位置和结束位置都不指定[:]
	// 从0下标开始，取值到下标2之前的下标，即，0，1
	var sce = arr1[0:2]
	// 切片len，结束位置-开始位置，2-0=2
	fmt.Println(sce, len(sce), cap(sce))
}

// 通过make方法
// 内部会先创建一个数组, 然后让切片指向数组
// 如果没有指定容量,那么容量和长度一样
func sliceCreateByMake() {
	// make(数据类型,长度,容量)
	var sce = make([]int, 3, 5)
	sce[2] = 123
	fmt.Println(sce, len(sce), cap(sce))
	/*
		内部实现原理
		var arr = [5]int{0, 0, 0}
		var sce = arr[0:3]
	*/
}

/*
通过Go提供的语法糖快速创建
和创建数组一模一样, 但是不能指定长度
通过该方式创建时, 切片的长度和容量相等
*/
func sliceCreateBySugar() {
	var sce = []int{1, 2}
	fmt.Println(sce, len(sce), cap(sce))
}

/*
添加数据, 扩容
添加数据时必须使用append方法
append函数会在切片末尾添加一个元素, 并返回一个追加数据之后的切片
如果追加之后没有超出切片的容量,那么返回原来的切片,
如果追加之后超出了切片的容量,那么返回一个新的切片
append函数每次给切片扩容都会按照原有切片容量*2的方式扩容
*/
func sliceAppend() {
	var sce = []int{1, 2}
	// 扩容为4
	sce = append(sce, 333)
	// 未超出,容量还是4
	sce = append(sce, 555)
	// 扩容为8
	sce = append(sce, 777)
	// 未超出,容量还是8
	sce = append(sce, 888)
	fmt.Println(sce, len(sce), cap(sce))
}

/*
拷贝
copy函数在拷贝数据时永远以小容量为准
*/
func sliceCopy() {
	var sce1 = []int{1, 3, 5}
	var sce2 = make([]int, 5)
	var sce3 = []int{6, 7, 9}
	var p1 *int = &sce1[0]
	var p2 *int = &sce2[0]
	var p3 *int = &sce3[0]
	fmt.Println(p1, sce1, p2, sce2)
	// sce1和sce2指向同一个数组
	// sce2和sce3指向不同数组
	sce2 = sce1
	p1 = &sce1[0]
	p2 = &sce2[0]
	fmt.Println(p1, sce1, p2, sce2)
	// copy(目标切片,源切片),将源切片数据copy导目标切片中
	copy(sce2, sce3)
	p1 = &sce1[0]
	p2 = &sce2[0]
	fmt.Println(p1, sce1, p2, sce2, p3, sce3)

	// 若切片起点相同,则指向同一个数组,修改一个元素会影响另一个元素
	arr1 := sce1[1:2]
	arr2 := sce1[1:2]
	var pa *int = &arr1[0]
	var pb *int = &arr2[0]
	arr2[0] = 12
	fmt.Println(pa, arr1, pb, arr2)
	// 切片只支持nil判断
	// fmt.Println(arr1==arr2) 编译报错
	fmt.Println(arr1 == nil)

	// 字符串的底层时[]byte数组,所以也支持切片
	strArr := "zhangsan"
	str1 := strArr[:2]
	str2 := make([]byte, 10)
	// 第二个参数只能是切片或者数组
	// copy(str2, []byte(strArr))
	copy(str2, strArr[2:4])
	fmt.Println(str1, str2)
}
