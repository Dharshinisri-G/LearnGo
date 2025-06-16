package main

import(
	"fmt"
)

func charandbyte(s string){
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n",rune,index)
	}
}

func printRunes(name string){
	fmt.Print("Characters:")
	runes:=[]rune(name)
	for i:= range runes {
		fmt.Printf("%c ",runes[i])
	}
	fmt.Println()
}


func printBytes(name string){
	fmt.Print("Bytes:")
	for i:= range name {
		fmt.Printf("%x ",name[i]) //x->hexadecimal
	}
	fmt.Println()
}

func printChars(name string){
	fmt.Print("Characters:")
	for i:= range name {
		fmt.Printf("%c ",name[i]) //c-> character
	}
	fmt.Println()
}

func main(){
	fmt.Println("Strings in Go")
	
	//String (slice of bytes in Go)
	name:="Sri" //enclosed in ""
	fmt.Println("String:",name) 

	//Accessing individual bytes of a string
	printBytes(name)

	//Accessing individual characters of a string
  printChars(name)
    // 🪲 example
  // sampstring:="Señor"
	sampstring:="Se\u00F1or"
	fmt.Println("String:",sampstring)
	printBytes(sampstring)
	printChars(sampstring)

	//Rune(Alias for int32 )
  sampStr:="Señor"
	fmt.Println("String2:",sampStr)
	printBytes(sampstring)
	printRunes(sampStr)

	//Accessing individual runes using for range loop
	charandbyte(sampStr)

}