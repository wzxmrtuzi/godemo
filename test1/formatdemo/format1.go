package formatdemo

import (
	"fmt"
	"strconv"
)

func Test() {
	str, error := strconv.ParseBool("t")
	if error == nil {
		fmt.Println(str)
	} else {
		fmt.Println(error.Error())
	}

	var preNum1 int32 = 198
	num := strconv.FormatInt(int64(preNum1), 10)
	fmt.Println(num)

	var preNum2 float32 = 198.234234
	// e 指数格式，f 小数格式
	num2 := strconv.FormatFloat(float64(preNum2), 'f', 6, 64)
	num3 := strconv.FormatFloat(float64(preNum2), 'e', 2, 64)
	fmt.Println(num2, num3)
}
