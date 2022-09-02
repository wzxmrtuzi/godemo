package reference

import (
	"fmt"
)

func Quote() {
	num := 123
	changeQuote(&num, 555)
	fmt.Println("quote", num)
}

func ArrTest() {
	arr := []int{1, 2, 3}
	changeArr(arr, 777)
	fmt.Println("ArrTest", arr)
}

// 在函数体内修改值类型参数, 不会影响到函数外的值
// func changeQuote(num, targeNum int) {
// 	num = targeNum
// }

func changeQuote(num *int, targeNum int) {
	*num = targeNum
}

func changeArr(arr []int, targeNum int) {
	arr[0] = targeNum
}
