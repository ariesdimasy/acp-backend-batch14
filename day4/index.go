package main

import (
	"acp-golang/day4/answer"
	"fmt"
)

func main() {
	fmt.Println("Hello Golang day 4")

	// soal 1
	// fmt.Println(answer.Compare("AKA", "AKASHI"))     // AKA
	// fmt.Println(answer.Compare("KANGOORO", "KANG"))  // KANG
	// fmt.Println(answer.Compare("KI", "KIJANG"))      // KI
	// fmt.Println(answer.Compare("KUPU-KUPU", "KUPU")) // KUPU
	// fmt.Println(answer.Compare("ILALANG", "ILA"))    // ILA

	// fmt.Println(answer.Caesar(3, "abc"))             // def
	// fmt.Println(answer.Caesar(2, "alta"))            // cnvc
	// fmt.Println(answer.Caesar(10, "alterraacademy")) // kvdobbkkmknowi
	// fmt.Println(answer.Caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	// // abcdefghijklmnopqrstuvwxyz
	// // bcdefghijklmnopqrstuvwxyza
	// fmt.Println(answer.Caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

	// fmt.Println(answer.ArrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}))   // [2 4]
	// fmt.Println(answer.ArrayUnique([]int{10, 20, 30, 40}, []int{5, 10, 15, 59})) // [20 30 40]
	// fmt.Println(answer.ArrayUnique([]int{1, 3, 7}, []int{1, 3, 5}))              // [7]
	// fmt.Println(answer.ArrayUnique([]int{3, 8}, []int{2, 8}))                    // [3]
	// fmt.Println(answer.ArrayUnique([]int{1, 2, 3}, []int{3, 2, 1}))              // []

	// fmt.Println(answer.FindMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2})) // 9
	// fmt.Println(answer.FindMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))    // 7
	// fmt.Println(answer.FindMaxSumSubArray(2, []int{2, 1, 4, 1, 1}))    // 5
	// fmt.Println(answer.FindMaxSumSubArray(3, []int{2, 1, 4, 1, 1}))    // 7
	// fmt.Println(answer.FindMaxSumSubArray(4, []int{2, 1, 4, 1, 1}))    // 8

	fmt.Println(answer.RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9})) // 4
	fmt.Println(answer.RemoveDuplicates([]int{2, 3, 4, 5, 6, 9, 9})) // 6
	fmt.Println(answer.RemoveDuplicates([]int{2, 2, 2, 11}))         // 2
	fmt.Println(answer.RemoveDuplicates([]int{2, 2, 2, 11}))         // 2
	fmt.Println(answer.RemoveDuplicates([]int{1, 2, 3, 11, 11}))     // 4

}
