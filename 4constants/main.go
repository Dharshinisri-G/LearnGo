package main

import(
	"fmt"
	"math"
)

func main(){
	fmt.Println("Constants in GO")

	//Declaring and Initializing a const
	const a=5
	// a=45 (can reassign value)
	fmt.Println("The value of a",a)

	//declare group of const
	const(
		retryLimit=4
		httpMethod="GET"
	)
	fmt.Println("Retry Limit",retryLimit)
	fmt.Println("http Method",httpMethod)

	//cannot be assigned to a value returned by a function call
	var a1=math.Sqrt(4)
	fmt.Println("a1 :",a1)
	// const a2=math.Sqrt(16) -> not allowed
  // fmt.Println("a2: ",a2)

	//String Constants & Typed and Untyped Constants
	const name="Sri" //Untyped Constants
	fmt.Println("Name :",name)
	var myName=name
	fmt.Println("My name is",myName)
  var defaultName="Dharshini"
	type myString string //myString is alias of string
	var customName myString="Sri"
	// defaultName=customName (defaultName of type string and another variable customName of type myString)
  fmt.Println("default:",defaultName,"Custom:",customName)

	//Boolean Constants (same as String Constants)
  const abool=true
	type myBool bool
	var bbool=false
	//abool=bbool ->>> not allowed abool type - Bool , bbool type myBool
	fmt.Println("bbool :",bbool)
	
	//Numeric Constants (int, float, complex)
	var i=5
	var f=5.4
	var c=5+2i
	fmt.Printf("Type of i %T, Type of f %T, Type of c %T\n",i,f,c)
    //untyped var changes to other type as needed
	const cc=5
	var intVar int=cc
	var int32Var int32=cc
	var float64Var float64=cc
	var complex64Var complex64=cc
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

	//Numeric Expressions
	var aVar=5+8.2
	fmt.Printf("The type of aVar is %T and the value is %v",aVar,aVar)
}
