package structdemo

import "fmt"

type User struct {
	name string
	age  int
}

func UserTest() {
	user := GetUser()
	SetAge(&user, 55)
	SetName(&user, "张三")
	fmt.Println("user", user)

	userInit := GetInitUser("里斯", 55)
	fmt.Println("userInit", userInit)

	userInit1()
	userInit2()

	convertOtherStruct()

	anonymousField()

	user.userInfo()
}

func userInit1() {
	type Annimal struct {
		name string
		age  int
		mp   map[int]string
	}
	// 初始化
	userVal := Annimal{name: "中"}
	userVal.age = 12
	mp := make(map[int]string)
	mp[2] = "李四"
	userVal.mp = mp
	fmt.Println(userVal)
}

func userInit2() {
	// 匿名结构体不用写类型,只能使用一次
	var person = struct {
		name string
		age  int
	}{name: "张三"}
	fmt.Println("anonymous-person", person)
}

func GetUser() User {
	return User{}
}

func GetInitUser(name string, age int) User {
	return User{name, age}
}

func SetName(user *User, name string) {
	user.name = name
}

func SetAge(user *User, age int) {
	user.age = age
}

// 类型转换
// 属性名、属性类型、属性个数、排列顺序都是类型组成部分
// 只有属性名、属性类型、属性个数、排列顺序都相同的结构体类型才能转换
func convertOtherStruct() {
	type Person1 struct {
		name string
		age  int
	}
	type Person2 struct {
		name string
		age  int
	}
	type Person3 struct {
		age  int
		name string
	}
	person1 := Person1{"zhangsan", 12}
	person2 := Person2(person1)
	// 编译报错
	// person3 := Person3(person1)
	// fmt.Println(person3)
	fmt.Println(person1)
	fmt.Println(person2)
}

// 匿名属性
// 没有指定属性名称,只有属性的类型, 我们称之为匿名属性
// 任何有命名的数据类型都可以作为匿名属性(int、float、bool、string、struct等)
func anonymousField() {
	type Person struct {
		int
		bool
		string
	}
	// 不指定名称初始化
	person1 := Person{2, true, "张三"}
	fmt.Println(person1)
	// 可以把数据类型作为名字显示初始化
	person2 := Person{int: 3, bool: false, string: "李四"}
	fmt.Println(person2)
	// 可以把数据类型当做属性名称操作结构体
	person2.int = 45
	fmt.Println(person2)

	// 结构体作匿名属性
	type User struct {
		Person
		string
	}
	user1 := User{Person{int: 3, bool: false, string: "李四"}, "张三"}
	// 如果要访问匿名结构体中的属性,可以先访问匿名结构体,再访问匿名结构体属性
	// 如果所有的匿名属性名称都不相同,可以不通过匿名结构体访问其属性
	fmt.Println(user1, user1.Person.string, user1.string)
}
func (user User) userInfo() {
	fmt.Println("userInfo", user.age, user.name)
}
