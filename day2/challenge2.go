package main

import (
	"fmt"
	"strconv"
)

func faktor_bilangan(num int32) {
	var i int32
	for i = 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}

}

func primeNumber(num int32) bool {
	var prime bool = true
	var i int32
	if num == 1 || num == 0 {
		prime = false
	}
	for i = 2; i < num; i++ {
		// println("--> num = ", num, " dan i = ", i)
		if num%i == 0 {
			prime = false

			//println("kokbisa num = ", num, " dan i = ", i)
		}
	}

	return prime
}

func palindrome(word string) bool {
	var isPalindrome bool = false
	var reverse string
	for i := len(word) - 1; i >= 0; i-- {
		reverse = reverse + string(word[i])
	}

	if reverse == word {
		isPalindrome = true
	}

	return isPalindrome
}

func pangkat(num int32, square int32) int32 {
	var result int32 = 1
	for i := 1; i <= int(square); i++ {
		result = result * num
	}

	return result
}

func FullPrima(num int32) bool {
	var isFullPrime bool = true
	var numStr string = strconv.FormatInt(int64(num), 10)

	for i := 0; i < len(numStr); i++ {
		isFullPrime = true
		newNum, err := strconv.ParseInt(string(numStr[i]), 10, 32)
		if err != nil {
			panic(err)
		}

		if primeNumber(int32(newNum)) == false {
			isFullPrime = false
			break
		}

	}

	//println(num, "==> ", isFullPrime)

	return isFullPrime
}

func playWithAsterik(n int) {
	var row string = ""
	var lengthRow int = n + (n - 1)
	var myCeil float64 = float64((lengthRow - 1) / 2)
	var middle int = int(myCeil)

	for i := 0; i < n; i++ {
		row = ""
		for j := 0; j < lengthRow; j++ {

			if (i == 0) && (j == middle) {
				row = row + "*"
			} else if j >= (middle-i) && j <= (middle+i) {
				if middle%2 == 0 {
					if (i % 2) == 0 {
						if j%2 == 0 {
							row = row + "*"
						} else {
							row = row + " "
						}

					} else {
						if j%2 == 0 {
							row = row + " "
						} else {
							row = row + "*"
						}
					}

				} else {
					if j == (middle-i) || j == (middle+i) {
						row = row + "*"
					} else if i%2 == 0 {
						if j%2 == 0 {
							row = row + " "
						} else {
							row = row + "*"
						}
					} else {
						if j%2 == 0 {
							row = row + "*"
						} else {
							row = row + " "
						}
					}
				}

			} else {
				row = row + " "
			}
		}

		println(row)
	}
}

func cetakTablePerkalian(num int32) {

	var row string
	var i int32
	var j int32
	for i = 1; i <= num; i++ {
		row = ""
		for j = i; j <= num*i; j = j + i {
			row = row + strconv.Itoa(int(j)) + " "
		}

		println(row)
	}
}

func ubahHuruf(sentence string) string {
	var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var newSentence string = ""

	for i := 0; i < len(sentence); i++ {

		if string(sentence[i]) == " " {
			newSentence += " "
			continue
		}

		findIndex := -1
		for a := 0; a < len(alphabet); a++ {
			if string(sentence[i]) == string(alphabet[a]) {
				findIndex = a
			}
		}

		if findIndex < 16 {
			findIndex += 10
			newSentence += string(alphabet[findIndex])
		} else {
			//println(string(alphabet[findIndex]), findIndex)
			findIndex = 9 - (25 - findIndex)
			newSentence += string(alphabet[findIndex])
		}

	}

	return newSentence

}

func main() {
	// faktor_bilangan(6)
	// faktor_bilangan(20)

	// fmt.Println(primeNumber(11)) // true
	// fmt.Println(primeNumber(13)) // true
	// fmt.Println(primeNumber(17)) // true
	// fmt.Println(primeNumber(20)) // false
	// fmt.Println(primeNumber(35)) // false

	// println("hahahaha")
	// fmt.Println(palindrome("civic"))       // true
	// fmt.Println(palindrome("katak"))       // true
	// fmt.Println(palindrome("kasur rusak")) // true
	// fmt.Println(palindrome("kupu-kupu"))   // false
	// fmt.Println(palindrome("lion"))        // false

	// fmt.Println(pangkat(2, 3))  // 8
	// fmt.Println(pangkat(5, 3))  // 125
	// fmt.Println(pangkat(10, 2)) // 100
	// fmt.Println(pangkat(2, 5))  // 32
	// fmt.Println(pangkat(7, 3))  // 343

	// fmt.Println(FullPrima(2))  // true
	// fmt.Println(FullPrima(3))  // true
	// fmt.Println(FullPrima(11)) // false
	// fmt.Println(FullPrima(13)) // false
	// fmt.Println(FullPrima(23)) // true
	// fmt.Println(FullPrima(29)) // false
	// fmt.Println(FullPrima(37)) // true
	// fmt.Println(FullPrima(41)) // false
	// fmt.Println(FullPrima(43)) // false
	// fmt.Println(FullPrima(53)) // true

	// playWithAsterik(3)
	// playWithAsterik(4)
	// playWithAsterik(5)
	// playWithAsterik(6)
	// playWithAsterik(7)

	//cetakTablePerkalian(9)

	fmt.Println(ubahHuruf("SEPULSA OKE"))     // COZEVCK YUO
	fmt.Println(ubahHuruf("ALTERRA ACADEMY")) // KVDOBBK KMKNOWI
	fmt.Println(ubahHuruf("INDONESIA"))       // SXNYXOCSK
	fmt.Println(ubahHuruf("GOLANG"))          // QYVKXQ
	fmt.Println(ubahHuruf("PROGRAMMER"))      // ZBYQBKWWOB

}
