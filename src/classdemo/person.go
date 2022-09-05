package classdemo

import (
	"fmt"
)

type PersonInterface interface {
	Eat()
}

type Person struct {
	name     string
	age      int
	identity string
}

type Student struct {
	Person
	grade string
	score float32
}

// 父类方法
func (person Person) Eat() {
	fmt.Println("person.eat", person.name)
}

// 子类重写父类方法
func (student Student) Eat() {
	fmt.Println("student.eat", student.name)
}

// 子类特有方法
func (student Student) GoToSchool() {
	fmt.Println("student go to school", student.name)
}

func (person *Person) SetAge(age int) {
	person.age = age
}

func (person *Person) SetName(name string) {
	person.name = name
}

func (person *Person) SetIdentity(identity string) {
	person.identity = identity
}

func (student *Student) SetPerson(person Person) {
	student.Person = person
}

func (student *Student) SetGrade(grade string) {
	student.grade = grade
}

func (student *Student) SetScore(score float32) {
	student.score = score
}
