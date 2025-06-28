package main

import(
	"fmt"
	"os"
)

func main(){
	fmt.Println("Error Handling in Go")
	f, err :=os.Open("/text.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}