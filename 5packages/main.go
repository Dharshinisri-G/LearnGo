package main

import (
	// _"5packages/simpleinterest" (if not used anywhere but wanna initialize)
	"5packages/simpleinterest"
	"fmt"
	"log"
)

var p, r, t = 500.0, 10.0, 1.0
// var p, r, t = -500.0, 10.0, 1.0


func init(){
	fmt.Println("Main function Initializing")
	if(p<0){
		log.Fatal("P is less than 0")
	}
	if(r<0){
		log.Fatal("r is less than 0")
	}
	if(t<0){
		log.Fatal("t is less than 0")
	}
}

var _ = simpleinterest.Calculate

func main(){
	
}

// func main(){
// 	fmt.Println("Packages in Go")
// 	//Packages are used to organize Go source code for better reusability and readability. Packages are a collection of Go sources files that reside in the same directory. Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.
// 	//Go Module
// 	   //A Go Module is nothing but a collection of Go packages. 

// 	fmt.Println("Simple interest calculation")

// 	//Importing the simpleinterest package in main
   
// 	 si:=simpleinterest.Calculate(p,r,t)
// 	 fmt.Println("SI :",si)

// 	 //A bit more on go build

// 	 //Exported Names
// 	   // Any variable or function which starts with a capital letter are exported names in go

// 	 //init function

//   //Use a blank identifier (error silencers)
// }