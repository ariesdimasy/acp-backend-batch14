package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Golang day 3")

	const gender string = "male"
	gender = "female"

	println("gender => ", gender)
}
