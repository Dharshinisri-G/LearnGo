package main

import(
	"fmt"
	"unsafe"
)

func main(){
	fmt.Println("Data Types in Go")
	
	//bool
	a,b :=true, false
	fmt.Println("a->",a,", b->",b);
	c:=a&&b
	fmt.Println("c->",c)
	d:=a||b
	fmt.Println("d->",d)

	//Signed Intergers int(32 or 64) -> int8, int16, int32, int64
	var a1 int=43
	b1:=65
	fmt.Println("a:",a1,",b:",b1)
  //finding type and size(in bytes)
  fmt.Printf("type of a is %T and the size of a is %d\n",a1,unsafe.Sizeof(a1))
  fmt.Printf("type of b is %T and the size of b is %d\n",b1,unsafe.Sizeof(b1))

	//Unsigned Integers(only positive +) uint(32 or 64) -> uint8, uint16, uint32, uint64
	var a2, b2 uint=45,54
  c2 :=a2*b2 //c2 is uint as a2 and b2 is uint
	fmt.Printf("The type of c2 is %T and the value of c2 is %d\n",c2,c2)

	//Floating Points
	a3, b3 :=4.53, 3.22
	fmt.Printf("Type of a: %T , b: %T\n",a3,b3)
	sum:=a3+b3
	diff:=a3-b3
	fmt.Printf("The sum of %f and %f is %f, and the difference is %f\n",a3,b3,sum,diff)
	a4,b4 :=4,3
	fmt.Printf("The sum of %d and %d is %d, and the difference is %d\n",a4,b4,a4+b4,a4-b4)

	//Complex Types -> complex64, complex128
	//func complex(r, i FloatType) ComplexType
	co1:=complex(5,7)
	co2:=3+21i
	sumCo:=co1+co2
	diffCo:=co1-co2
	fmt.Println("The sum of",co1,"and",co2,"is",sumCo,"and the difference is",diffCo)

// 	Other numeric types
// byte is an alias of uint8
// rune is an alias of int32

  //String Type
	firstName:="Dharshini"
	lastName:="Sri"
	fullName:=firstName+" "+lastName;
  fmt.Println("My name is",fullName)

	//Type Conversion
	at:=3
	bt:=4.2
	sumT:=at+int(bt)
	fmt.Println("The sum of",at,"and",bt,"is",sumT)
	//assign a variable of one type to another using Type Conversion
	i:=20
	var j float64 =float64(i)
	fmt.Println("j =",j)

}