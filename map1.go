// //Q - 4.2 find the duplicates of the slice and return them in  the map - done

package main

import "fmt"

func main() {
	a := []int{}
	b := 0
	fmt.Println("Enter the numbers for the slice. Note := if you enter '0' loop will break")
	for true {
		fmt.Scan(&b)
		if b == 0 {
			break
		}
		a = append(a, b)
	}
	fmt.Println("The duplicate items in the slice are", map1(a))
}

func map1(a []int) map[int]int {
	b := make(map[int]int)
	c := make(map[int]int)
	for _, val := range a {
		b[val]++
	}
	for key, value := range b {
		if value > 1 {
			c[key] = value
		}
	}
	return c
}

// func main() {
// 	 a := map[int]int{
// 		1:4 ,
// 		2:5 ,
// 		3:6 ,
// 		9:5 }
// 	 for key,value := range a {
// 		fmt.Println(key,value)
// 	}
// }
