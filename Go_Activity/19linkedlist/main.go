package main

import (
	"fmt"
)

var head, tail *node

type node struct{
	data int
	next *node
}

func addElement(newNode *node){
	if head==nil{
		head=newNode
		tail=newNode
	}else{
		tail.next=newNode
		tail=newNode
	}
}

func printElement(){
	current:=head
	for current!=nil{
		fmt.Print(current.data)
		if current.next!=nil { fmt.Print("->")}
		current=current.next
	}
}

func removeElementByValue(value int){
	current:=head
	var prev,nextt *node
	nextt=current.next
  for current.data != value{
      prev=current
			current=current.next
      nextt=current.next
	}
	if(prev!=nil){
  prev.next=nextt
	return
	}
  head=head.next
}

func removeElementByIndex(index int){
	 currIndex:=0
	 current:=head
	var prev,nextt *node
	nextt=current.next
  for currIndex!=index{
      prev=current
			current=current.next
      nextt=current.next
			currIndex++
	}
	if(prev!=nil){
  prev.next=nextt
	return
	}
	head=head.next
}

func main(){
  fmt.Println("Linked List in Go. Add , remove and list the elements.");
  addElement(&node{5,nil})
  addElement(&node{10,nil})
  addElement(&node{15,nil})
  addElement(&node{20,nil})
  addElement(&node{25,nil})
	removeElementByValue(20)
	removeElementByIndex(1)
	printElement();
}