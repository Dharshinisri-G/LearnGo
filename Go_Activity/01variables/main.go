package main

import(
  "fmt"
)

func main(){
  fmt.Println("Ways to declare Variables in Go")
	var a int = 5
	fmt.Println("The value of a is",a)
  b:=6
	fmt.Println("The value of b is",b)
	var c=3
	fmt.Println("The value of c is",c)
	var d,e int = 4,7
	//var d,e = 4,7
	//d,e:=4,7
	fmt.Println("The value of d and e are",d,"and",e)
	var f int
	f=8
	fmt.Println("The value of f is",f);
	const PI=3.14
	fmt.Println("The value of PI is",PI)
	var(
		firstname = "Dharshini"
		lastname = "Sri"
		age = 21
	)
	fmt.Println(firstname,lastname,"is",age,"years old")
}