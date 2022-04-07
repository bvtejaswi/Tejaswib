package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter Array 1 Size: ")
	fmt.Scanln(&n)
	a1 := make([]int, n)
	fmt.Println("Enter Array Elements ")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter %d Element: ", i+1)
		fmt.Scanln(&a1[i])
	}
	fmt.Print("Enter Array 2 Size: ")
	fmt.Scanln(&n)
	a2 := make([]int, n)
	fmt.Println("Enter Array Elements")
	for i := 0; i < n; i++ {
		fmt.Printf("Enter %d Element: ", i+1)
		fmt.Scanln(&a2[i])
	}
	comparison(a1, a2)
}

func comparison(a1, a2 []int) {
	s1 := len(a1)
	s2 := len(a2)
	// Storing second array elemets - so i can check with arr1 elements is it present in arr2 or not
	myMap := map[int]int{}
	for i := 0; i < s2; i++ {
		myMap[a2[i]]++
	}
	for i := 0; i < s1; i++ {
		_, ok := myMap[a1[i]]
		if !ok {
			fmt.Printf("This Number %d is Not present in Second Array\n", a1[i])
		}
	}
}
