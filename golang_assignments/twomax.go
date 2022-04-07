package main

import (
	"fmt"
	"math"
)

func twoLargest(arr []int, size int) {
	if size <= 1 {
		return
	} else {
		var first int = math.MinInt64
		var second int = math.MinInt64
		for i := 0; i < size; i++ {
			if first < arr[i] {
				second = first
				first = arr[i]
			} else if second < arr[i] {
				second = arr[i]
			}
		}
		fmt.Println("First Largest   : ", first)
		fmt.Println("Second Largest  : ", second)
	}
}
func main() {
	var arr = []int{
		1,
		3,
		33,
		8,
		18,
		9,
		4,
		5,
		7,
	}
	var size int = len(arr)
	twoLargest(arr, size)
}
