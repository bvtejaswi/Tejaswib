package main

import "fmt"

type Stack []string

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

func (st *Stack) Push(str string) {
	*st = append(*st, str)
}

func (st *Stack) Pop() bool {
	if st.IsEmpty() {
		return false
	} else {
		index := len(*st) - 1
		*st = (*st)[:index]
		return true
	}
}

func (st *Stack) Top() string {
	if st.IsEmpty() {
		return ""
	} else {
		index := len(*st) - 1
		element := (*st)[index]
		return element
	}
}

func prec(s string) int {
	if s == "^" {
		return 3
	} else if (s == "/") || (s == "*") {
		return 2
	} else if (s == "+") || (s == "-") {
		return 1
	} else {
		return -1
	}
}

func main() {

	infix := "2+4*(2^3-2)^(2+5*2)-4"
	postfix := infixToPostfix(infix)
	fmt.Printf("%s infix ", infix)
	fmt.Printf("converted to %s postfix ", postfix)

}

func infixToPostfix(infix string) string {
	var sta Stack
	var postfix string
	for _, char := range infix {
		opchar := string(char)
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			postfix = postfix + opchar
		} else if char == '(' {
			sta.Push(opchar)
		} else if char == ')' {
			for sta.Top() != "(" {
				postfix = postfix + sta.Top()
				sta.Pop()
			}
			sta.Pop()
		} else {
			for !sta.IsEmpty() && prec(opchar) <= prec(sta.Top()) {
				postfix = postfix + sta.Top()
				sta.Pop()
			}
			sta.Push(opchar)
		}
	}
	for !sta.IsEmpty() {
		postfix = postfix + sta.Top()
		sta.Pop()
	}
	return postfix
}
