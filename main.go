package main

import (
	"demo1/test1/calculate"
	"demo1/test1/enumdemo"
	"demo1/test1/formatdemo"
	"demo1/test1/functiondemo"
	"demo1/test1/judge"
	"demo1/test1/reference"
	"demo1/test1/stringutil"
	"demo1/test1/structdemo"
	"fmt"
	"strconv"
	"unsafe"
)

// 常量
const AGE = 12
const (
	PI         = 3.14
	HOUR       = 24
	HOUR2      // 和上一行的有值的变量相同
	HOUR3      // 和上一行的有值的变量相同
	num1, num2 = 1, 2
	num3, num4 // 和上一行的有值的变量相同,且个数相同
)

func main() {
	res := calculate.Sum(2, 4)
	fmt.Println("res=", res)

	// var name = "zhang"

	// var name
	// name = "zhang"
	// 变量组
	var (
		name1 = "123"
		name2 = "456"
	)
	name := "zhang"
	name = "lisi"
	name1 = "aaa"
	name2 = "bbb"
	age := 45
	resStr := stringutil.Append(name, "san")
	fmt.Println("name", resStr, "size", unsafe.Sizeof(resStr), stringutil.Append(resStr, "123"), name1, name2, age, AGE)
	fmt.Println((3.3 / 4))

	formatdemo.Test()

	enumdemo.EnumTest()

	user := structdemo.GetUser()
	fmt.Println(structdemo.GetName(user), structdemo.GetAge(user))
	user = structdemo.SetAge(user, 55)
	fmt.Println(user)

	fmt.Println(strconv.FormatInt(int64(19), 2))

	// judge.JudgeTest()

	judge.ForTest()

	sum, avg := functiondemo.FuncTest()
	fmt.Println("123", sum, avg)

	reference.Quote()

	reference.ArrTest()
}
