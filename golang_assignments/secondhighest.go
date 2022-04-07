package main

import "fmt"

func main() {
	arr := []int{9, 2, 7, 1, 3}
	var first int = 0
	var second int = 0

	first = arr[0]
	for i := 1; i <= 4; i++ {
		if first < arr[i] {
			second = first
			first = arr[i]
		} else if second < arr[i] {
			second = arr[i]
		}
	}
	fmt.Println("Second largest element is: ", second)
}
