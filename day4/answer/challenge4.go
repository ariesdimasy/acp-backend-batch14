package answer

func Compare(a, b string) string {
	// your code here
	var firstCompare string = a
	var haystack string = b
	var result = ""

	if len(a) >= len(b) {
		firstCompare = b
		haystack = a
	}

	for i := 0; i < len(haystack); i++ {
		var temp string = ""
		for j := 0; j < len(firstCompare); j++ {
			temp += string(haystack[j])
		}

		if temp == firstCompare {
			result = temp
			break
		}
	}

	return result

}

// masih error
func Caesar(offset int, input string) string {
	var alphabet string = "abcdefghijklmnopqrstuvwwxyz"
	var newSentence string = ""

	for i := 0; i < len(input); i++ {

		if string(input[i]) == " " {
			newSentence += " "
			continue
		}

		findIndex := -1
		for a := 0; a < len(alphabet); a++ {
			if string(input[i]) == string(alphabet[a]) {
				findIndex = a
			}
		}

		if findIndex < (26 - offset) {
			findIndex += offset

		} else {
			//println(string(alphabet[findIndex]), findIndex)
			findIndex = (offset - 1)

		}

		newSentence += string(alphabet[findIndex])

	}

	return newSentence

}

func ArrayUnique(arrayA, arrayB []int) []int {

	var result = []int{}
	for i := 0; i < len(arrayA); i++ {

		var isSame bool = false
		for j := 0; j < len(arrayB); j++ {
			if arrayA[i] == arrayB[j] {
				isSame = true
				break
			}

		}

		if !isSame {

			result = append(result, arrayA[i])
		}

	}

	return result

}

func FindMaxSumSubArray(k int, arr []int) int {
	// your code here

	return 0
}
