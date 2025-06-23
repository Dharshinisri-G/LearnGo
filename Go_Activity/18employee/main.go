package main

import(
  "fmt"
)
type Employee struct{
	Name string
	Id int 
	Salary int
}
type behaviour interface{
	display();
	giveRise(amount int)
}
func (e Employee) display(){
	fmt.Println("Name: ",e.Name);
	fmt.Println("Id: ",e.Id);
	fmt.Println("Salary: ",e.Salary);
}
func (e *Employee) giveRise(amt int){
	e.Salary+=amt
}
func main(){
  fmt.Println("Rectangle struct with fields for width and height.method to calculate the area of the rectangle.");
  var emp1 behaviour=&Employee{
		"Dharshinisri",
		21,
		23000,
	}
	emp1.display();
	var amount int
	fmt.Scanln(&amount)
  emp1.giveRise(amount)
	emp1.display()
}