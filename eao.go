// q-4.1 -Even or odd numbers 
package main

import "fmt"

func eao(s []int)([]int,[]int){
	b:= []int{}
	c:= []int{}
	for i := range s{
		if(s[i]%2==0){
			b=append(b,s[i])
		}else{
			c=append(c,s[i])
		}
	}
	return b,c
}

func main() {
	a := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	b,c := eao(a)
	fmt.Println("The even numbers of the slice are",b)
	fmt.Println("The odd numbers of the slice are",c)
}
