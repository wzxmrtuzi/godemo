package main

import (
	"fmt"
	"godemo/src/arrdemo"
	"godemo/src/calculate"
	"godemo/src/enumdemo"
	"godemo/src/formatdemo"
	"godemo/src/functiondemo"
	"godemo/src/judge"
	"godemo/src/reference"
	"godemo/src/slicedemo"
	"godemo/src/stringutil"
	"godemo/src/structdemo"
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

	reference.MapTest()

	sum, _ = functiondemo.Anonymous1()
	_, avg = functiondemo.Anonymous1()
	fmt.Println(sum, avg)
	sum, _ = functiondemo.Anonymous2()
	fmt.Println(sum)

	functiondemo.TodayTodo()

	functiondemo.ClosePackTest()

	functiondemo.DelayTest()

	functiondemo.InitTest()

	arrdemo.ArrayTest()

	slicedemo.SliceTest()
}
