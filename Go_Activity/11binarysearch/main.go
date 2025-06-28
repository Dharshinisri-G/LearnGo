package main

import(
  "fmt"
)

func main(){
  fmt.Println("Binary search using for loop on a sorted slice");
	a:=[]int{2,4,6,8,10, 13, 16, 20}
	var t int
	fmt.Scanf("%d",&t)
	l,h:=0,len(a)-1
	m:=(l+h)/2
	var index int
	for l<=h {
		if(a[m]==t){
       index=m
			 break
		}else if(a[m]>t){
       h=m-1
		}else{
			l=m+1
		}
		m=(l+h)/2
	}
	fmt.Println("The target element",t,"is at the index",index)
}