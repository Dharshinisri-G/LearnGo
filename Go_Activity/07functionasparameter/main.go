package main

import(
  "fmt"
)
func hello(){
	fmt.Println("Hello! Welcome to the Party");
}
func bye(function func()){
	function();
	fmt.Println("Bye! Catch you later.");
}
func main(){
  fmt.Println("Passing Function as parameter");
  bye(hello);
}