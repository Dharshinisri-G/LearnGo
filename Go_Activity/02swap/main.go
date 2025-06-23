package main

import(
  "fmt"
)

func main(){
  fmt.Println("Swapping variable without temp variable");
	a,b:=2,3
	a,b=b,a
	fmt.Println("a :",a,", b :",b);

}