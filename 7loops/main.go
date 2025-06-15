package main

import(
	"fmt"
)

func main(){
	fmt.Println("Loops in Go")

	//For loop Syntax
	//for initialisation; condition; post {}
     for i:=1; i<=10; i++ {
			fmt.Println("hi",i)
		 }
	
	//Break
	  for i:=1; i<=10; i++ {
			if i>5 {
				break
			}
			fmt.Println("Hii",i)
		}	 
		fmt.Println("Loop ended")

		//Continue
		 for i:=1; i<=10; i++ {
			if i%2!=0 {
				continue;
			}
			fmt.Println("Value",i,)
		 }
		 fmt.Println("The even numbers are displayed")

		 //Nested for loop
		 for i:=1; i<=5; i++{
			for j:=1; j<=i; j++{
				fmt.Print("*")
			}
			fmt.Println()
		 }

		 //Labels (break the outer loop from inside the inner for loop)
		 for i:=1; i<=3; i++ {
			for j:=1; j<=3; j++{
				fmt.Println("i =",i,", j =",j)
			}
		 }
		  //Wanna break the inner and outer when i and j are equal
		 outer:	
		 for i:=0; i<=3; i++ {
			for j:=1; j<=3; j++{
				fmt.Println("i =",i,", j =",j)
				if i==j {
					break outer
				}
			}
		 }

		 //Implement While using the for loop
		  //  initialisation and post are omitted
		i:=1
		for ;i<=10; { //for i<=10 {} (semi-colons can be ommited)
			fmt.Println(i)
			i+=2
		}

		//Multiple Variable Declaration
		for no,i := 10,1; i<=10 && no<=20; no,i = no+1,i+1 {
			fmt.Printf("%d * %d = %d\n",no,i,no*i)
		}

		//Infinite Loop for{}
		for{
			fmt.Println("Hello Loop")
		}
}