package main

import "fmt"

func main() {

	var i int
	rev := 0

	fmt.Print("Enter Number: ")
	var num int
	fmt.Scanln(&num)

	for temp := num; temp > 0; temp = temp / 10 {
		i = temp % 10
		rev = rev*10 + i
	}

	if num == rev {
		fmt.Println(num, " is a Palindrome")
	} else {
		fmt.Println(num, " is Not a Palindrome")
	}
}
