package main

import(
	"fmt"
)

func main(){
	fmt.Println("If else Statements in Go")

	//If Statement syntax 
	  //if condition {}  the braces { } are not optional
	a:=6
	if a%2==0{
		fmt.Println("a is even")
		// return
	}
	fmt.Println("a is odd")

	//if else Statement
	b:=7
	if b%2==0{
		fmt.Println("b is evern")
	} else{
		fmt.Println("b is odd")
	}

  //if else if else Statement
	passengerAge:=10
	ticketprice:=0
	if passengerAge < 5{
		ticketprice=0
	} else if passengerAge < 22{
    ticketprice=10
	} else{
		ticketprice=15
	}
  fmt.Println("The price of the ticket is",ticketprice)

	//If with Assignment (Short Hand Assignment)
	//if assignment-statement; condition {}
	  if age:=10; age<11{          // age is limited to the if block
			fmt.Println("Age is",age)
		}

		//Gotcha
		  //else statement should start in the same line after the closing curly brace } of the if statement

		//Idiotic Go (end if with return if you don't need to use else)
			
}