package functiondemo

import "fmt"

/**
匿名函数场景：
1. 当某个函数只需要被调用一次时
2. 需要执行一些不确定的操作时
*/
/**
闭包：
闭包是一个特殊的匿名函数, 它是匿名函数和相关引用环境组成的一个整体
也就是说只要匿名函数中用到了外界的变量, 那么这个匿名函数就是一个闭包
*/
func ClosePackTest() {
	num := 10
	aFunc := func() {
		// 闭包中使用的变量和外界的变量是同一个变量, 所以可以闭包中可以修改外界变量
		num = 6
		fmt.Println("afunc", num)
	}
	fmt.Println("execute before", num)
	aFunc()
	fmt.Println("execute after", num)
	// 得到一个闭包
	// 只要闭包还在使用外界的变量, 那么外界的变量就会一直存在
	resIncrease := increase()
	fmt.Println(resIncrease())
	fmt.Println(resIncrease())
	fmt.Println(resIncrease())
}

func increase() func() int {
	num := 1
	return func() int {
		// 匿名函数中用到了addUpper中的x,所以这是一个闭包
		num++
		return num
	}
}

func FuncTest() (sum int, avg int) {
	sum = 12
	avg = 2
	// 直接调用,匿名函数
	func(s string) {
		fmt.Println(s)
	}("a ba")

	// 保存到变量
	strFunc1 := func(s string) string {
		return s
	}
	str1 := strFunc1("a str1")
	fmt.Println("str1", str1)

	// 作为参数
	parameterFunc(func(s string) {
		fmt.Println(s)
	})

	// 作为返回值函数
	returnValue := returnValueFunc()
	returnValue(4, 5)
	return
}

// 函数作为参数
func parameterFunc(f func(s string)) {
	f("hello parameter_func")
}

// 函数作为返回值
func returnValueFunc() func(int, int) {
	return func(a int, b int) {
		fmt.Println(a + b)
	}
}

// 全局匿名函数方式1，定义在函数外的匿名函数：全局匿名函数
var Anonymous1 = func() (sum int, avg int) {
	sum = 12
	avg = 2
	return
}

// 全局匿名函数方式2
var (
	Anonymous2 = func() (sum int, avg int) {
		sum = 12222
		avg = 2
		return
	}
)

func TodayTodo() {
	zhangsan := func() {
		fmt.Println("上班")
	}
	lisi := func() {
		fmt.Println("睡午觉")
	}
	today(zhangsan, "zhangsan")
	today(lisi, "lisi")
}

func today(custom func(), name string) {
	fmt.Print(name)
	morning := func() {
		fmt.Print("起床")
		fmt.Print("吃早餐")
	}
	evening := func() {
		fmt.Print("回家")
		fmt.Println("吃晚餐")
	}
	morning()
	custom()
	evening()
}
