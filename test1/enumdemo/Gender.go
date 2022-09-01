package enumdemo

import (
	"fmt"
)

func EnumTest() {
	// 在同一个常量组中,只要上一行出现了iota,那么后续行就会自动递增
	// iota从0开始
	// iota也支持常量组+多重赋值, 在同一行的iota值相同
	const (
		MALE            = iota
		FEMALE, UNKNOWN = iota, iota
		A, B
	)
	fmt.Println(MALE, FEMALE, UNKNOWN, A, B)
	fmt.Printf("aaa = %T\n", A)
}
