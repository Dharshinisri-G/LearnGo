package main

import (
	"fmt"
)

//function declaration
//   func functionname(parametername datatype) returntype {
//  //function body
// }
// parameters and return type are optional

func calculateBill(price int,quantity int) int{ 
	//consecutive parameters are of the same type, parameters can be written as (price, quantity int)
	var total=price*quantity
	return total
}

//Multiple return values 
		func rectProps(length, width float64)(float64, float64){
			area:=length*width
			perimeter:=(length+width)*2
			return area,perimeter
		}

//Named Return values (named values are in the first line of the function)
   func rectProps2(length, width float64)(area, perimeter float64){
		area=length*width;
		perimeter=(length+width)*2
		return
	 }

func main(){
	fmt.Println("Functions in Go")
	var price,quantity=24,2
  totalPrice:=calculateBill(price,quantity)
	fmt.Println("totalPrice :",totalPrice)
	var length, width =2.1, 1.1
	area,perimeter:=rectProps(length,width)
	fmt.Println("Area :",area,", perimeter :",perimeter)
//Blank Identifier (to discard any place or type)
  area2,_:=rectProps2(length,width)
	fmt.Println("The area of Rect2",area2)
}



