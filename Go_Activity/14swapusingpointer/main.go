package main

import(
  "fmt"
)
func swap(ap *int,bp *int){
	*ap,*bp=*bp,*ap
}
func main(){
  fmt.Println("Swapping values of two integer using pointer");
	a,b:=4,5
	ap:=&a
	bp:=&b
	swap(ap,bp)
  fmt.Println("a :",a,", b :",b);
}