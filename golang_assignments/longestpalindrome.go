package main

import "fmt"

func printSubStr(str string, low int, high int) {
	for i := low; i <= high; i++ {
		fmt.Println(str[i])
	}
}

func longestPalSubstr(str string) {
	n := len(str)
	maxLength := 1
	start := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			flag := 1

			for k := 0; k < (j-i+1)/2; k++ {
				if str[i+k] != str[j-k] {
					flag = 0
				}
			}

			if flag != 0 && (j-i+1) > maxLength {
				start = i
				maxLength = j - i + 1
			}

		}
	}
	//fmt.Println("Longest palindrome subString is:", str, start, start+maxLength-1)
	fmt.Println("Longest palindrome subString is:", str[start:start+maxLength-1])

}

func main() {

	fmt.Print("Enter string: ")
	var str string
	fmt.Scanln(&str)
	longestPalSubstr(str)

}
