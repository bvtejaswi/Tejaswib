package main

import (
	"fmt"
)

func second(arr []int) {
	first := arr[0]
	second := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > first {
			second = first
			first = arr[i]
		} else if arr[i] > second && arr[i] < first {
			second = arr[i]

		}
	}
	if second != -1 {
		fmt.Printf("second max: %d\n", second)

	} else {
		fmt.Println("No second max")
	}

}

func main() {
	fmt.Print("Enter array size: ")
	var n int
	fmt.Scanln(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter %d element: ", i+1)
		fmt.Scanln(&array[i])
	}
	second(array)

}
