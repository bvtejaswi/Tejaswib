package main

import "fmt"

type bag struct {
	bat  int
	ball int
}

func (f bag) countForBall() {
	fmt.Printf("In a bag there are %v bats & %v balls\n", f.bat, f.ball)
}

type total struct {
	bag
}

func main() {
	var f = bag{1, 2}
	f.countForBall()
	var apple = total{bag{4, 1}}
	apple.countForBall()
}
