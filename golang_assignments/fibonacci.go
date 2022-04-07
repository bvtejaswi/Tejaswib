package main

import "fmt"

func main() {
	var len int
	var next int
	first := 0
	second := 1

	fmt.Print("Enter the count of numbers: ")
	fmt.Scanln(&len)

	fmt.Printf("%v terms of fib \n", len)
	for i := 0; i < len; i++ {
		if i <= 1 {
			next = i
		} else {
			next = first + second
			first = second
			second = next
		}
		fmt.Printf("next term is %v \n", next)

	}

}
