package main

import(
	"fmt"
)

func change(s ...string){
	s[0]="Go"
}

func change2(s ...string){
	s=append(s,"playground")
	fmt.Println(s)
}

func findd(num int, nums []int) {
	fmt.Printf("Type of nums is %T\n",nums)
	found:=false
	for i,v := range nums {
		if v==num {
       fmt.Println(num,"is found at index",i,"in",nums)
			 found=true
		}
	}
	if !found {
		fmt.Println(num,"not found in",nums)
	}
	fmt.Println()
}

func find(num int, nums ...int) {
	fmt.Printf("Type of nums is %T\n",nums)
	found:=false
	for i,v := range nums {
		if v==num {
       fmt.Println(num,"is found at index",i,"in",nums)
			 found=true
		}
	}
	if !found {
		fmt.Println(num,"not found in",nums)
	}
	fmt.Println()
}

func main(){
	fmt.Println("Variadic Functions in Go")

	//variable number of arguments func hello(a int, b ...int) {}
	//hello(1, 2) //passing one argument "2" to b
// hello(5, 6, 7, 8, 9) //passing arguments "6, 7, 8 and 9" to b
  find(1,1,3,4)
	find(3,4,5,3,7,8)
	find(9,10,11)
	find(12)

  //Slice Arguments
	fmt.Println("Using Slice Arguments")
	findd(1,[]int{1,3,4})
	findd(3,[]int{4,5,3,7,8})
	findd(9,[]int{10,11})
	findd(12,[]int{})
 
	//Append is a variadic function
	//func append(slice []Type, elems ...Type) []Type

	//Passing a slice to a variadic function
	nums:=[]int{13,141,15,16}
	// find(24,nums) ---> Errror (cannot use nums (type []int) as type int in argument to find)
  find(24,nums...)     //Slice Expansion | Syntatic Sugar

	//Gotcha
	welcome:=[]string{"hello","world"}
	change(welcome...) //the actual slice is passed
	fmt.Println(welcome) //changes happen in the slice

  change2(welcome...)
}