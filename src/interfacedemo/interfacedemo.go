package interfacedemo

import (
	"fmt"
)

// 接口中只能有方法的声明不能有方法的实现
// 接口中只能有方法,不能有字段
// 只有实现了接口中所有的方法, 才算实现了接口, 才能用该接口类型接收,只实现部分接口会报错
// 接口中嵌入接口时不能出现相同的方法名称

// 超集
type annimal interface {
	eat()
	run()
}

// 子集
type action interface {
	annimal
	breathe()
}

// 空接口,可以接收任意类型数据, i = 1,i="ff"...
var i interface{}

// 自定义类型
type str string

// 自定义类型实现接口
func (s str) eat() {
	fmt.Println("custom", s)
}

func (s str) run() {
	fmt.Println("custom", s)
}

type Dog struct {
	foot  string
	mouth string
}
type Bird struct {
	foot  string
	mouth string
}

func InterFaceTest() {
	dog := Dog{"四只脚", "一个嘴"}
	dog.eat()
	dog.run()
	bird := Bird{"一对翅膀", "一个尖嘴"}
	bird.run()
	bird.eat()

	var s str = "ffff"
	s.run()

	// 接口类型变量,只能调用实现的方法,无法访问其属性
	var annimalInterFace annimal = Dog{"四只脚", "一个嘴"}
	annimalInterFace.eat()
	// 接口类型变量,将接口类型还原为原始类型,才能访问属性
	if d, ok := annimalInterFace.(Dog); ok {
		fmt.Println("ok-idiom模式转换", d)
	}
	switch d := annimalInterFace.(type) {
	case Dog:
		fmt.Println("switch模式转换", d)
	default:
		fmt.Println("无此类型")
	}

	// 将抽象接口类型转换为具体接口类型
	// 定义一个抽象接口
	var i interface{}
	i = Dog{"四只脚", "一个嘴"}
	// i还不能调用eat方法和run方法,因为接口没有抽象接口没有定义
	// 需要转为具体接口
	if s, ok := i.(annimal); ok {
		s.eat()
	}
}

func (dog Dog) eat() {
	fmt.Println("吃", dog.mouth)
}

func (dog Dog) run() {
	fmt.Println("跑", dog.foot)
}

func (bird Bird) eat() {
	fmt.Println("尖嘴", bird.mouth)
}
func (bird Bird) run() {
	fmt.Println("翅膀", bird.foot)
}
