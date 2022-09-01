package structdemo

// import "fmt"

type User struct {
	name string
	age  int
}

func GetUser() User {
	return User{}
}

func GetName(user User) string {
	return user.name
}

func SetName(user User, name string) {
	user.name = name
}

func GetAge(user User) int {
	return user.age
}

func SetAge(user User, age int) User {
	user.age = age
	return user
}
