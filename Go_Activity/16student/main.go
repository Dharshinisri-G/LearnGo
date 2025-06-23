package main

import(
  "fmt"
)
type Person struct{
	name string
	age int
}
type Student struct{
	Person
	major string
}
func (s Student) studentDetails(){
	fmt.Println("Name:",s.name);
	fmt.Println("Major:",s.major);
	fmt.Println("Age:",s.age);
}
func main(){
  fmt.Println("Student struct that embeds the Person struct and adds a field for the major, display student details");
	student1:=Student{
		Person: Person {
		name:"Dharshinisri",
		age:21,
	},
		major: "Information Technology",
}
	student1.studentDetails();
}