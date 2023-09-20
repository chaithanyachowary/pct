package main

import (
	"fmt"
	"time"
)

// func a(b, c int) int {
// 	return c + b
// }

// type x func(int, int) int

// func main() {

// 	var y x
// 	y=a
// 	a := y(2, 2)
// 	fmt.Println(a)

// }

// struct and function
// type emp struct{
// firstname string
// lastname string
// age int
// salary
// IsFullTime bool
// }
// type salary struct{
// 	bonus int
// 	stock int
// 	base int
// }
// func main() {
// 	a := emp{
// 		firstname : "John",
// lastname : "John",
// age : 30,
// salary:salary{
// 	bonus : 1,
// 	stock : 1,
// 	base : 1,
// },
// IsFullTime : true,
// }
// 	c(a)
// }

// func c(b emp){
// 	fmt.Println(b.firstname, b.lastname,b.bonus)
// }

// stuct,func and slices
// type book struct {
// 	bookName string
// 	Author   string
// 	year     int
// }

// func main() {
// 	books := []book{
// 		{"12 Rules for a life", "Jordan B Peterson", 2012},
// 		{"Atomic Habbits", "James Clear", 2018},
// 		{"Immortals of Meluha", "Amish Tripati", 2010},
// 		{"The Oath Of the vayuputras", "Amish Tripati", 2013},
// 	}
// 	newbook := book{"The Secret of Nagase", "Amish Tripati", 2011}

// 	books = append(books, newbook)
// 	modify(books)
// 	for _, book := range books {
// 		fmt.Printf("Book name :'%s';authors : '%s' and published year: '%d' '\n'", book.bookName, book.Author, book.year)
// 	}
// }

// func modify(a []book) {
// for b, _ := range a {
// 	a[b].year += 2
// 	fmt.Println(a[b].year)
// }
// for i := 0; i < len(a); i++ {
// 	a[i].year += 2
// }
//fmt.Println(a)
// }

// func main(){
//  a := book{"I die,my die, why do you care!","priya darshi",2016}
// //c := 12
// b(&a)
// fmt.Println(a)
// }

// func b(a *book){
//  a.year += 2
// 	//book+=1
// 	fmt.Println(a)
// }

// // map and struct
// type book struct {
// 	name   string
// 	author string
// 	year   int
// }

// func main() {
// 	books := make(map[book]int)
// 	books[book{"12 Rules for a life", "Jordan B Peterson", 2012}] = 1
// 	books[book{"13 Rules for a life", "Jordan B Peterson", 2014}] = 2
// 	for key, value := range books {
// 		fmt.Println(key.year)
// 		fmt.Println(key.author)
// 		fmt.Println(key.name)
// 		fmt.Println(value)
// 	}

// }

// leet code :- two sum problem

// func twoSum(a []int, target int) []int {
//     b :=[]int{}
//     sum :=0
//     for i:=0;i<len(a)-1;i++{
//         for j:=i+1;j<len(a);j++ {
//             sum=a[i]+a[j]
//             if sum==target{
// 				b =append(b,i,j)
//                 break
//             }else{
//                 sum=0
//             }
//         }
//     }
// 	return b
// }

// func main() {
// 	a:=[]int{1,2,3,4,5}
// 	b:= 8
// 	fmt.Println(twoSum(a,b))
// }

// leet code := Parentesis
// func isValid(s string) bool {
// 	if len(s)%2 != 0 {
// 		 return false
// 	}
//     for i:= 0; i<=len(s)-2; i=i+2{
// 		if s[i] ==91  && s[i+1] != 93 || s[i] == 40 && s[i+1] != 41 ||s[i] == 123 && s[i+1] != 125{
// 			return false}
// 		}
// return true
// }

// func main() {
// 	var s string
// 	fmt.Scan(&s)
// 	fmt.Println(isValid(s))
// }

// leet code := max profit
// func main() {
// 	a:= [] int{7,3,5,8,1,6,4,2}
// 	fmt.Println(max(a))
// }

// func max(a []int) int {
// 	b:= min(a)

// 	return max(b)
// }

//leet code := palindrome
// func isPalindrome(x int) bool {
//     i:= x
// 	sum:=0
// 		for i>0 {
// 			sum=(sum*10)+(i%10)
// 			i=i/10
// 		}
// 		fmt.Println(sum)
// 		if sum==x{
// 			return true
// 		}else{
// 			return false
// 		}
// }
// func main() {
// 	a:=121
// 	fmt.Println(isPalindrome(a))
// }
// leet code : duplicate
// func main() {
// 	a:=[]int{1,2,3,4,5}
// 	fmt.Println(containsDuplicate(a))
// }
// func containsDuplicate(nums []int) bool {
//     for i:=0;i<len(nums)-1;i++ {
// 		for j:=i+1;j<len(nums);j++ {
// 			if nums[i]==nums[j]{
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// leet code :-
// func reverseString(s []byte)  {
// 	b:=[]byte
//     for i:=len(s);i>=0;i--{
// 			b =append(b,s[i])
// 		}
// }
// func main() {
// 	s:= []byte{"h","e","l","l","o"}

// }

// llet code : square root
// func main() {
// 	a:=2
// 	fmt.Println(sqrt(a))
// }

// func sqrt(x int) int {
// 	i:=1
// 	if x ==0{
// 		return x
// 	}
// 	for i<x/2 {
// 		if x/(i*i)==1&&x%(i*i)==0{
// 			break
// 		}
// 		if x<=i*i {
// 			fmt.Println(i)
// 			i=i-1
// 			fmt.Println(i)
// 			break
// 		}
// 		i++
// 	}
// 	fmt.Println(i)
// 	return i
// }
