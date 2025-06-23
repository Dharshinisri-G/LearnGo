package main

import(
  "fmt"
)
type Person struct{
	name string
	age int
}
func (p Person) personDetails(){
	fmt.Println("Name:",p.name);
	fmt.Println("Age:",p.age);
}
func main(){
  fmt.Println("Person struct with fields for name and age.Creating a method to display the person's details.");
	person1:=Person{
		name:"Dharshinisri",
		age:21,
	}
	person1.personDetails();
}