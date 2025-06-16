package main

import (
	"fmt"
	"reflect"
)

func mapEquals(map1, map2 map[string]int) bool{
	if len(map1)!=len(map2) {
		return false
	}
	for k,v1 := range map1 {
		v2, ok := map2[k] 
		 if !ok || v1!=v2 {
			return false
		 }
	}
	return true
}

func main(){
	fmt.Println("Maps in Go")

	//How to create a map?  (make(map[type of key]type of value))
	currencycode:=make(map[string]string)
	fmt.Println(currencycode)

	// Adding items to a map (order is not preserved)
	currencycode["INR"]="Indian Rupee"
	currencycode["EUR"]="Euro"
	currencycode["USD"]="United States Dollar"
	fmt.Println(currencycode)

	//initialize a map during the declaration itself
	currencycode2:= map[string]string {
		"USD":"United States Dollar",
		"INR":"Indian Rupee",
		"EUR":"Euro",
	}
	currencycode2["GBP"]="Pound Streling"
	fmt.Println(currencycode2)

	//nil map panics
	//var nilmap map[string]string
	//nilmap["USD"]="United States Dollar" ----> panic: assignment to entry in nil map

	//Retrieving value for a key from a map
	fmt.Println("The currency name for the INR is",currencycode["INR"])
  fmt.Println("The currency name for the GBP is",currencycode["GBP"]) //empty string if key not present
	currency := "USD"
	currencyName := currencycode[currency]
	fmt.Println("Currency name for currency code", currency, "is", currencyName)

	//Checking if a key exists (value, ok := map[key])
	cucode:="GBP"
	if currencyname, ok := currencycode[cucode]; ok{
		fmt.Println("The currencyname of",cucode,"is",currencyname)
	} else{
		fmt.Println("The",cucode,"is not present in the map")
	}
  
	//Iterate over all elements in a map
	for code, name := range currencycode2 {
		fmt.Println("The currencyname of",code,"is",name)
	}

	//Deleting items from a map
	fmt.Println("Before Deletion",currencycode2)
	delete(currencycode2,"GBP")
	fmt.Println("After Deletion",currencycode2)

	//Map of structs
  type curr struct{
		name string
		symbol string
	}
	curINR:=curr{
		name:"Indian Rupee",
		symbol:"₹",
	}
	curUSD:=curr{
		name:"United States Dollar",
		symbol:"$",
	}
	curEUR:=curr{
		name:"Euro",
		symbol:"€",
	}
	currencyCode:= map[string]curr {
		"USD":curUSD,
		"INR":curINR,
		"EUR":curEUR,
	}
	for cyCode, cyInfo := range currencyCode {
		fmt.Println("The Currency name",cyInfo.name,"and the symbol",cyInfo.symbol,"of currency",cyCode)
	}

	//Length of the map
	fmt.Println("The length of the Map currencyCode is",len(currencyCode))

	//Maps are reference types
	employeeSalary:= map[string]int {
		"John":23000,
		"Mike":24500,
		"Sara":45000,
	}
	fmt.Println("Original Employee Salary",employeeSalary)
	modified:=employeeSalary
	modified["Mike"]=43000
	fmt.Println("Modified Employee Salary",employeeSalary)

	//Maps equality
	//Maps can’t be compared using the == operator. The == can be only used to check if a map is nil.
	map1:= map[string]int{
		"one":1,
		"two":2,
	}
	map2:=map1
	fmt.Println(map2)
	// if map2==map1{
     //---> Error: invalid operation: map1 == map2 (map can only be compared to nil)
	// }

	//Using reflection to compare two maps have same key-value pairs
	if reflect.DeepEqual(map1,map2) {
		fmt.Println("the map1 and map2 are equal")
	} else{
		fmt.Println("The map1 and map2 are not equal")
	}

	//Other way to compare two maps
	if mapEquals(map1, map2){
		fmt.Println("Maps are equals")
	} else {
		fmt.Println("Maps are not equals")
	}
}