package src

import (
	"math"
)

func PrimeNumber(number int) bool {

	if 0 < number && number < 3 {
		return false
	} else if number%2 == 0 {
		return false
	}

	for i := 3; i < int(math.Floor(float64(number/2))); i = i + 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func Pow(x, n int) int {

	var result int = 1
	var once bool = false

	if n%2 == 1 {
		once = true
	}

	for i := 1; i <= int(math.Floor(float64(n/2))); i++ {
		result = result * (x * x)
	}

	if once {
		result = result * x
	}

	return result
}

func JoinArrayRemoveDuplicate(arrayA, arrayB []string) []string {

	var newArr = []string{}

	for i := 0; i < len(arrayA); i++ {

	}

	// your code here
	return []string{"", ""}
}
