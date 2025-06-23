package main

import (
	"08functioncall/employee"
	"fmt"
)

func main(){
  fmt.Println("Two different package and make a function call between it");
	employee.EmployeeDetails("Sri",005)
}