package main

import(
  "fmt"
)

func main(){
  fmt.Println("Data Types in Go");
	var a int = 45 //integer
	var b float64 = 54.32 //float
  var c string = "Sri" //string
	var d bool = false //boolean
	var f complex64 = 43+2i  //Complex
	var g uint = 67   //unsigned integer
	var h byte = 32  //byte (uint8)
	var i rune = 456 //rune (int32)
	fmt.Println("Integer: ",a);
	fmt.Println("Unsigned Integer: ",g);
	fmt.Println("Float: ",b);
	fmt.Println("String: ",c);
	fmt.Println("Boolean: ",d);
	fmt.Println("Complex: ",f);
	fmt.Println("Byte: ",h);
	fmt.Println("Rune: ",i);
}