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
	var arr2 [3]int = [3]int{1, 2, 3}
	changeArr2(arr2, 777)
	fmt.Println("ArrTest", arr)
	fmt.Println("ArrTest", arr2)
}

// 在函数体内修改值类型参数, 不会影响到函数外的值
// func changeQuote(num, targeNum int) {
// 	num = targeNum
// }

func changeQuote(num *int, targeNum int) {
	*num = targeNum
}

// 影响函数外的值
func changeArr(arr []int, targeNum int) {
	arr[0] = targeNum
}

// 不影响函数外的值
func changeArr2(arr [3]int, targeNum int) {
	arr[0] = targeNum
}

func MapTest() {
	mp := map[string]string{"name": "zhangsan", "age": "12"}
	changeMap(mp, "里斯")
	fmt.Println("map", mp)
}

// 影响函数外的值
func changeMap(mp map[string]string, targeStr string) {
	mp["name"] = targeStr
}
