package errordemo

import (
	"errors"
	"fmt"
)

func ErrorTest() {
	createError1()
	createError2()
	res, err := divide(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	res2 := div(3, 0)
	fmt.Println(res2)
}

func createError1() {
	err1 := fmt.Errorf("错误信息")
	fmt.Println("err1", err1)
}

func createError2() {
	var err error = errors.New("错误信息2")
	fmt.Println("err2", err.Error())
}

func divide(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
	} else {
		res = a / b
	}
	return
}

// panic会中断程序
// 多个异常,只有第一个会被捕获
// 如果有异常写在defer中, 那么只有defer中的异常会被捕获
func div(a, b int) (res int) {
	// 定义一个延迟调用的函数,用于捕获panic异常
	// 要在panic前定义,可以在调用div方法前使用
	defer func() {
		if err := recover(); err != nil {
			res = -1
			fmt.Println(err)
		}
	}()
	if b == 0 {
		panic("除数不能为0")
	} else {
		res = a / b
	}
	return
}
