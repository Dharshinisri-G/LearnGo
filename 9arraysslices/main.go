package main

import (
	"fmt"
)

func countriess()[]string{
	countries:=[]string{"India","Russia","China","korea","Brazil","England"}
	countriesneeded:=countries[:len(countries)-2]
	countriesCpy:=make([]string,len(countriesneeded))
	copy(countriesCpy,countriesneeded)
	return countriesCpy
	
}

func substractOne(numbers []int){
	for i:=range numbers{
		numbers[i]-=2
	}
}

func printarray(arr [3][2]string){
	for _,v1:=range arr{
		for _,v2:=range v1{
			fmt.Print(v2," ")
		}
		fmt.Println()
	}
}

func changeArray(arr1 [5]int){
 arr1[0]=33
 fmt.Println("Inside function:",arr1)
}

func main(){
	fmt.Println("Arrays and Slices in Go")

	//Ways of Declaration
	var a [3]int
	fmt.Println(a)

	//Index from 0 to length-1
	a[0]=0
	a[1]=1
	a[2]=2
	fmt.Println(a)
	
	//Short hand declaration
	b:=[3]int{3,4,5}
	fmt.Println(b)

	//not necessary that all elements in an array have to be assigned a value during short hand declaration
	c:=[4]int{6}
	fmt.Println(c)

	// ... makes the compiler determine the length
	d:=[...]int{7,8,9}
	fmt.Println(d)

	//not possible to assing arrays to another with different distinct types
	// c=d  --> error

	//Arrays are value types (not reference types)
  countries:=[...]string{"USA","China","Korea","India","Brazil"}
  countriescopy:=countries
	countriescopy[0]="Singapore"
	fmt.Println("countries",countries)
	fmt.Println("countries",countriescopy)

	//arrays are passed to functions as parameters, they are passed by value and the original array is unchanged
	arr1:=[...]int{1,2,3,4,5}
	fmt.Println("before",arr1)
	changeArray(arr1)
	fmt.Println("After",arr1)

	//Length of an array
	ar1:=[...]int{6,7,8,9,10}
	fmt.Println("The length of the array",ar1,"is",len(ar1))

	//Iterating arrays using for loop
  for i:=0; i<len(ar1); i++ {
		fmt.Println("The value at index",i,"is",ar1[i])
	}

	//range form of the for loop (include index and value)
	sum:=int(0)
	for i,v :=range ar1 {  // _,v :=range ar1 (to ignore index)
		fmt.Println("The element at",i,"is",v)
		sum+=v
	}
	fmt.Println("Sum of Array ar1 is",sum)

	//Multi-dimensional Array
	a2:=[3][2]string{
		{"lion","tiger"},
		{"cat","dog"},
		{"penquin","parrot"},
	}
	printarray(a2);
	var b2 [3][2]string
	b2[0][0]="Dosa"
	b2[0][1]="Idly"
	b2[1][0]="Samosa"
	b2[1][1]="Puffs"
	b2[2][0]="Fried Rice"
	b2[2][1]="Chicken"
	printarray(b2);

	                                                     //Slices
 //Creating a slice
  as:= [4]int{1,2,3,4} 
	var bs []int = as[1:3] //slice from start to end-1
	fmt.Println(bs)

	//creates an array and returns a slice reference
	cs:=[]int{3,4,5}
	fmt.Println(cs)

	//Modifying a slice
	darr:=[...]int{1,2,3,4,5,6,7,8,9,10}
	dslice:=darr[2:5]
	fmt.Println("Array Before",darr)
	for i:=0; i<len(dslice); i++ {
		//for i:=range dslice {
		dslice[i]++
	}
	fmt.Println("Array After",darr)

	//creates a slice which contains all elements of the array
	numa:=[3]int{23,43,54}
	num1:=numa[:]
	num2:=numa[:]
	fmt.Println("Initial Array:",numa)
	num1[0]=534
	fmt.Println("Array after change in slice num1",numa)
	num2[1]=678
	fmt.Println("Array after change in slice num1",numa)
    //slices share the same array.
	
	//length(elements in slice) and capacity(elements in array from the start of slice) of a slice
	fruits:=[7]string{"apple","banana","orange","grapes","pomegranate","strawberry","dragonfruit"}
	fruitssliced:=fruits[1:3]
	fmt.Printf("The length of the slice: %d and the capacity of the slice: %d\n",len(fruitssliced),cap(fruitssliced))
   
	//A slice can be re-sliced upto its capacity
	fruitssliced=fruitssliced[:cap(fruitssliced)]
	fmt.Printf("The length of the slice: %d and the capacity of the slice: %d\n",len(fruitssliced),cap(fruitssliced))

	//creating a slice using make (func make([]T, len, cap))
	imake:=make([]int,5) //([]int,5,5)
	fmt.Println(imake)

	//Appending to a slice (func append(s []T, x ...T) []T.)
  cars:=[]string{"Lambo","Ferrari","Porsche"}
	fmt.Println("Cars:",cars,"has old length",len(cars),"and the capacity",cap(cars))
	cars=append(cars,"BMW")
	fmt.Println("Cars:",cars,"has new length",len(cars),"and the capacity",cap(cars))

	//zero value of a slice is nil
  var names[] string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names=append(names, "john","Merry","Henna")
    fmt.Println("names contents:",names)
	}

	//append one slice to another using the ... operator
  veggies:=[]string{"Carrot","Brinjal"}
	fruitss:=[]string{"apple","Mango","Orange"}
	food:=append(veggies, fruitss...)
	fmt.Println(food)

	//Passing a slice to a function
   numbers:=[]int{8,9,7}
	 fmt.Println("Before Calling Func",numbers)
	 substractOne(numbers)
	 fmt.Println("After Calling Func",numbers)

	 //Multidimensional slices
	 prgms:=[][]string {
		{"C","Java"},
		{"Javascript"},
		{"Python","Go"},
	 }
	 for _,v1:=range prgms{
		for _,v2:=range v1{
			fmt.Print(v2," ")
		}
		fmt.Println()
	 }

	 //Memory Optimization
	 countriesNeeded:= countriess()
	 fmt.Println(countriesNeeded)


}

