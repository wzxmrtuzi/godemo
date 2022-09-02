package judge

import (
	"fmt"
)

func JudgeTest() {
	var arr1 [3]int = [3]int{1, 1, 3}
	var p *int = &arr1[0]
	fmt.Println(arr1, p)
	if arr1[1] == 2 {
		fmt.Println("123", arr1[1])
	}
	if age := 18; age >= 18 {
		fmt.Println("adulthood")
	} else {
		fmt.Println("children")
	}

	var num = 19
	// case可以放函数，变量，常量，表达式
	switch getNumber(12) {
	case 18, getNumber(19):
		fmt.Println("adulthood")
	case num:
		fmt.Println("graduate")
	default:
		value := "unknown"
		fmt.Println(value)
	}
}

func ForTest() {
	var arr1 [3]int = [3]int{1, 1, 3}
	fmt.Println(arr1)
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(arr1[i])
	// }
	var ran [2]int = [2]int{2, 1}
	for i, v := range ran {
		fmt.Println(i, v, arr1[i])
	}
}

func getNumber(val int) int {
	return val
}
