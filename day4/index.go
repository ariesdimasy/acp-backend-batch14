package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Golang day 4")

	var colors = []string{"red", "yellow", "blue", "brown"}

	colors = append(colors, "orange")

	var someColors = make([]int, 10)
	// println(colors, someColors)

	fmt.Println(colors, someColors)

	multi := [][]int{} // ini slice
	multi = append(multi, []int{1, 2})
	multi = append(multi, []int{3, 4})

	fmt.Println((multi))
}
