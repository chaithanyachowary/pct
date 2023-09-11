package main

import "fmt"






// Q-4.4 find the character repetition number
func main(){
	// a,b :="this is sample s s string","s"
	var a string
	var b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	c:=0
	for i:=0;i<len(a);i++{
		if a[i]==b[0] {
			c++	}
	}
	fmt.Printf("The middle index of %s in string is %d.",b,c/2)
}	
