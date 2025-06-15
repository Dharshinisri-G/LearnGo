package main

import(
	"fmt"
	"math"
)

func main(){
	//Variable Declaration
	var myAge int 
	fmt.Println("My Intial age is",myAge)
	myAge=12
	fmt.Println("My age after 1st assignment",myAge);
	myAge=21
	fmt.Println("My age after 1st assignment",myAge);
  fmt.Println()

  //Variable Declaration with initial Value
	var sisAge int=17
	fmt.Println("Sister's Intial age is",sisAge)
  fmt.Println()	

	//Type Inference
	var broAge=12
	fmt.Println("Brother's Age is",broAge)
  fmt.Println()

	//Multiple Variable Declaration
	var myName, sisName, broName string="Sri","Prithi","Naren" //type String can be removed as the value is initialized
	fmt.Println("My Name:",myName,", Sister's Name:",sisName,", Brother's Name",broName)
  fmt.Println()

	//declaring multiple variables with type inference
	var price, quantity = 100, 5
	fmt.Println("The price is",price,"and the quantity is",quantity)

	//Different Data types in a single Statement
	var(
		name="Dharshinisri"
		age=21
		height int
	)
	fmt.Println("My name is",name,"\nMy age is",age,"\nMy height before initializing",height)

	//Short hand declaration (Must initialize value)
	count:=19
	fmt.Println("The count is",count)

	//Multiple variable declaration using short hand declaration
	Name, Age := "Dharshini",21
	// name, age := "Naveen" error
	fmt.Println("my name is",Name,"and my age is",Age)

	//SHD for already initialized variable
	a,b :=10,20
	fmt.Println("a=",a,", b=",b)
	b,c :=30,40
	fmt.Println("a=",a,", b=",b,", c=",c)
	// b,c :=50,60 --> Error (No new variables)
	b,c =50,60
	fmt.Println("a=",a,", b=",b,", c=",c)

  //Assigning Values to the variable during run time
	aa,bb :=32.22,23.11
	cc :=math.Min(aa,bb)
	fmt.Println("The Minimum of a and b is",cc)

	//variables declared as belonging to one type cannot be assigned a value of another type
	// agee:=45
	// agee="Gaja"
	// Error (Untyped String Constant)

}