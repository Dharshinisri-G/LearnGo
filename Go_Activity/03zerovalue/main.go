package main

import(
  "fmt"
)

func main(){
  fmt.Println("Zero Value in Go");
	fmt.Println("Zero value is the default value to a variable when it is declared but no values are assigned to it");
	fmt.Println("Zero value for different types");
	var a int  //integer
	var b float64 //float
  var c string //string
	var d bool  //boolean
	var e chan int //Channel
	var f complex64   //Complex
	var g uint    //unsigned integer
	var h byte   //byte (uint8)
	var i rune //rune (int32)
	var j *int //pointer
	var k [4]int //Array
	var l []float64 //Slice
	var m map[string]string //Map
	var n interface{} //Interface
	fmt.Println("Integer: ",a);
	fmt.Println("Float: ",b);
	fmt.Println("String: ",c);
	fmt.Println("Boolean: ",d);
	fmt.Println("Complex: ",f);
	fmt.Println("Unsigned Integer: ",g);
	fmt.Println("Byte: ",h);
	fmt.Println("Rune: ",i);
	fmt.Println("Channel: ",e);
	fmt.Println("Pointer: ",j);
	fmt.Println("Array: ",k);
	fmt.Println("Slice: ",l);
	fmt.Println("Map: ",m);
	fmt.Println("Interface: ",n);
}