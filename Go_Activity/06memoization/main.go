package main

import(
	"fmt"
)

var a []int

func fibo(n int) int{
	if n<=1{
		a[n]=n
		return a[n]
	}
	if a[n]!=0{
		return a[n]
	}
  a[n]=fibo(n-1)+fibo(n-2)
	return a[n]
}

func main(){
	fmt.Println("Fibonacci");
	var n int
	fmt.Scanf("%d",&n);
	a=make([]int,n+1)
	for i := 0; i < n; i++{
      fmt.Print(fibo(i)," ");
	}
}