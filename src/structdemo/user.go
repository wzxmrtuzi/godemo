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
	fmt.Println("user", user, GetName(&user), GetAge(&user))

	userInit := GetInitUser("里斯", 55)
	fmt.Println("userInit", userInit)

	userInit1()
	userInit2()

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
	// 匿名结构体不用写类型
	var person = struct {
		name string
	}{name: "张三"}
	fmt.Println("anonymous-person", person)
}

func GetUser() User {
	return User{}
}
func GetInitUser(name string, age int) User {
	return User{name, age}
}

func GetName(user *User) string {
	return user.name
}

func SetName(user *User, name string) {
	user.name = name
}

func GetAge(user *User) int {
	return user.age
}

func SetAge(user *User, age int) {
	user.age = age
}
