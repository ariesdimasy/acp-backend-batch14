package main

import (
	"acp-golang/day5/answer"
	"fmt"
)

func main() {
	fmt.Println("Hello Golang day 5")

	// a := 10
	// b := 20

	// answer.Swap(&a, &b)
	// fmt.Println(a, b)

	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = answer.MinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("nilai min => ", min)
	fmt.Println("nilai max => ", max)

}
