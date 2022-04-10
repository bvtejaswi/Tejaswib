package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

//push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)

}

//pop

func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toPop := s.items[l]
	s.items = s.items[:l]
	return toPop

}

func main() {
	var str string
	fmt.Print("Enter string: ")
	fmt.Scanln(&str)
	if isBalanced(str) {
		fmt.Printf("%s is Balanced\n", str)
	} else {
		fmt.Printf("%s is not Balanced\n", str)

	}
}

func isMatching(p int, q int) bool {
	if p == '(' && q == ')' {
		return true
	} else if p == '{' && q == '}' {
		return true
	} else if p == '[' && q == ']' {
		return true
	} else {
		return false
	}
}

func isBalanced(str string) bool {
	s := Stack{}
	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == '{' || str[i] == '[' {
			s.Push(int(str[i]))
		}
		if str[i] == ')' || str[i] == '}' || str[i] == ']' {
			if s.IsEmpty() {
				return false
			} else if !(isMatching(s.Pop(), int(str[i]))) {
				return false

			}

		}

	}
	return s.IsEmpty()
}
