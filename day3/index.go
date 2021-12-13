package main

import (
	"acp-golang/day3/src"
	"fmt"
)

func main() {
	fmt.Println("Hello Golang day 3")

	// fmt.Println(src.PrimeNumber(1000000007))  // true
	// fmt.Println(src.PrimeNumber(1500450271))  // true
	// fmt.Println(src.PrimeNumber(1000000000))  // false
	// fmt.Println(src.PrimeNumber(10000000019)) // true
	// fmt.Println(src.PrimeNumber(10000000033)) // true

	// fmt.Println(src.Pow(2, 3))  // 8
	// fmt.Println(src.Pow(7, 2))  // 49
	// fmt.Println(src.Pow(10, 5)) // 100000
	// fmt.Println(src.Pow(17, 6)) // 24137569
	// fmt.Println(src.Pow(5, 3))  // 125

	// Test cases
	// fmt.Println(src.JoinArrayRemoveDuplicate([]string{"apel", "anggur"}, []string{"lemon", "leci", "nanas"}))
	// // ["apel", "anggur", "lemon", "leci", "nanas"]

	// fmt.Println(src.JoinArrayRemoveDuplicate([]string{"samsung", "apple"}, []string{"apple", "sony", "xiaomi"}))
	// // ["samsung", "apple", "sony", "xiaomi"]

	// fmt.Println(src.JoinArrayRemoveDuplicate([]string{"football", "basketball"}, []string{"basketball", "football"}))
	// // [football basketball]

	// fmt.Println(src.MunculSekali("1234123"))    // [4]
	// fmt.Println(src.MunculSekali("76523752"))   // [6, 3]
	// fmt.Println(src.MunculSekali("12345"))      // [1 2 3 4 5]
	// fmt.Println(src.MunculSekali("1122334455")) // []
	// fmt.Println(src.MunculSekali("0872504"))    // [8 7 2 5 4]

	fmt.Println(src.PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(src.PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(src.PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(src.PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(src.PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]

}
