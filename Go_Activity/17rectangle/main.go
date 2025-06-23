package main

import(
  "fmt"
)
type Rectangle struct{
	width float64
	height float64
}
func (r Rectangle) area(){
	fmt.Printf("The area of the rectangle is %.1f",(r.height*r.width))
}
func main(){
  fmt.Println("Rectangle struct with fields for width and height.method to calculate the area of the rectangle.");
  rect1:=Rectangle{
		width: 2.3,
		height: 1.2,
	}
	rect1.area();
}