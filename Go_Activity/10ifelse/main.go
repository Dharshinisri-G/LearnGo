package main

import(
  "fmt"
)

func main(){
  fmt.Println("If else Statement in Go");
	//Syntax if(condition){}else{}
	var a int
	fmt.Scanln(&a)
	if(a%2==0){
		fmt.Println(a,"is even");
	}else{
		fmt.Println(a,"is odd");
	}
}