// // Q-4.3 find the string is a palindrome or not
package main

import "fmt"

func main() {
	// fmt.Println("Enter a string")
	// var s string
	// fmt.Scan(&s)

	s := "gfchvjbkn"
	fmt.Println(palindrome(s))

}
func palindrome(s string) bool {
	//chaithanya
	b := len(s) / 2
	a := 0
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			break
		}
		if s[i] == s[j] {
			a++
		}
	}
	if a == b {
		return true
	} else {
		return false
	}
}
