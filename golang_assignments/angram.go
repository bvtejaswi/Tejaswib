package main

import "fmt"

func isAnagram(str1 string, str2 string) bool {

	lenS := len(str1)
	lenT := len(str2)

	if lenS != lenT {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {
		anagramMap[string(str1[i])]++
	}

	for i := 0; i < lenT; i++ {
		anagramMap[string(str2[i])]--
	}

	for i := 0; i < lenS; i++ {
		if anagramMap[string(str1[i])] != 0 {
			return false
		}
	}

	return true
}

func main() {
	var str1 string
	var str2 string
	fmt.Println("Enter First String")
	fmt.Scanf("%s", &str1)
	fmt.Println("Enter Second String")
	fmt.Scanf("%s", &str2)

	if isAnagram(str1, str2) {
		fmt.Printf("%s and %s are anagram\n", str1, str2)
	} else {
		fmt.Printf("%s and %s are not anagram\n", str1, str2)
	}
}
