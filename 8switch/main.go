package main

import(
	"fmt"
	"math/rand"
)

func number() (num int){
	num=15*3
	return
}

func main(){
	fmt.Println("Swtich Statements in Go")

	//Swtich (idiomatic way of replacing complex if else clauses)
	finger:=2
	fmt.Print("Finger ",finger," is ")
	switch finger{ //switch finger:=8; finger{} (finger is limited to the switch block)
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	// case 4:
	// 	fmt.Println("Another Ring")  ----> Duplicate cases are not allowed
	case 5:
		fmt.Println("Pinky")
  default: //default case when all the above is false
		fmt.Println("Invalid Input")
	}

	//Multiexpressions in case
		letter := "i"
		switch letter {
		case "a", "e", "i", "o", "u":
			fmt.Println(letter,"is a vowel")
		default:
			fmt.Println(letter,"is a Consonant")
		}

		//Expressionless switch
		hour:=15
		switch {
		case hour>6 && hour<=12 :
        fmt.Println("Morning Shift")
		case hour>12 && hour<=18:
		    fmt.Println("Afternoon Shift")
		case hour>18 && hour<=22:
			  fmt.Println("Evening Shift")
		default:
			 fmt.Println("Invalid Shift hour")
		}

		//fallthrough
		switch num :=number(); {
		case num<50:
			fmt.Println(num,"less than 50")
			fallthrough
		case num<100:
			fmt.Println(num,"less than 100")
			fallthrough
		case num<200:
			fmt.Println(num,"less than 200")
			fallthrough
		case num<25:
			fmt.Println(num,"is less than 25")
			fallthrough
			//Note: ---->Fallthrough happens even when the case evaluates to false
		default:
			fmt.Println(num,"greater than all other numbers in this case")
		}

		//Breaking switch
		switch num:=-5; {
		case num<50:
			fmt.Println("Entered Switch statement")
			if num<0 {
				break;
			}
			fmt.Println(num,"less than 50")
			fallthrough
		case num<100:
			fmt.Println(num,"less than 100")
			fallthrough
		case num<200:
			fmt.Println(num,"less than 200")
		}

		//Breaking the outer for loop
		randloop:
		for{
			switch i:=rand.Intn(100); {
			case i%2==0:
				fmt.Println("gen even number is",i)
				break randloop;
			}
		}

		
}