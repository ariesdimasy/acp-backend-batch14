package src

import (
	"math"
	"strconv"
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
	var combine = append(arrayA, arrayB...)

	for i := 0; i < len(combine); i++ {
		var exists bool = false
		for j := 0; j < len(newArr); j++ {
			if combine[i] == newArr[j] {
				exists = true
				break
			}
		}

		if !exists {
			newArr = append(newArr, combine[i])
		}
	}

	// your code here
	return newArr
}

func MunculSekali(angka string) []int {
	// your code here
	var result = []int{}

	for i := 0; i < len(angka); i++ {
		var misuketa bool = false
		for j := 0; j < len(angka); j++ {
			//fmt.Println(string(angka[i]), string(angka[j]), "\t")
			if string(angka[i]) == string(angka[j]) && i != j {
				//fmt.Println("MISUKETAZOO")
				misuketa = true
				break
			}
		}

		if !misuketa {
			strAngka := string(angka[i])
			newNum, err := strconv.ParseInt(strAngka, 10, 32)
			if err != nil {
				panic(err)
			}
			result = append(result, int(newNum))
		}
	}

	return result
}

func PairSum(arr []int, target int) []int {

}
