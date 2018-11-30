package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

type Student struct {
	Person
	Id    string
	Grade string
	Name  string
}

type Teacher struct {
	*Person //指针类型，使用前先初始化
	Course  string
	Name    string
}

//假修改
func VChangeStu(s Student) {
	//因为Person不是指针，所以没有修改
	s.Name = "Change Name"
	s.Person.Name = "Change Person Name"
}

//真修改
func RChangeStu(s *Student) {
	s.Name = "Change Name"
	s.Age = 13
	s.Person.Name = "Change Person Name" //可以修改
}

func RChangeTea(t *Teacher) {
	t.Name = "我是老师"
	t.Sex = "男"
	t.Person.Name = "我是人民教师"
}

func main() {
	s1 := Student{Person: Person{Name: "张三", Age: 13, Sex: "男"}, Id: "13321", Grade: "三年级"}
	fmt.Printf("%+v\n", s1)

	s2 := Student{Person{"张三", 13, "男"}, "12312", "三年级", "Stu张三"}
	fmt.Printf("%+v\n", s2)

	var s3 Student
	s3.Name = "s3Name"
	s3.Person.Name = "S3Person Name"
	fmt.Printf("%+v\n", s3)

	VChangeStu(s3)
	fmt.Printf("%+v\n", s3)

	RChangeStu(&s3)
	fmt.Printf("%+v\n", s3)

	var t1 Teacher
	t1.Name = "T1 Name"
	t1.Person = new(Person)
	t1.Person.Name = "T1 Person Name"
	fmt.Printf("%+v\n", t1.Person.Name)

	var t2 Teacher
	t2.Name = "T2 Name"
	t2.Person = &Person{}
	t2.Person.Name = "T2 Person Name"
	fmt.Printf("%+v\n", t2.Person.Name)

	var t3 Teacher
	t3.Person = new(Person)

	RChangeTea(&t3)
	fmt.Printf("%+v,%+v\n", t3, t3.Person)

}
