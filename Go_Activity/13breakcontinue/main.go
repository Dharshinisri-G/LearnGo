package main

import(
  "fmt"
)

func main(){
  fmt.Println("Break and Continue in Go");
	//Even numbers 
	for i:=0;i<=10;i++{
      if(i%2!=0){
				continue;
			}
			fmt.Println(i);
			if(i==8){
          break;
			}
	}
}