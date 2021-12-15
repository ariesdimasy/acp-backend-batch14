package answer

func MinMax(numbers ...*int) (min int, max int) {
	max = *numbers[0]
	min = *numbers[0]
	mynums := numbers

	for i := 1; i < len(mynums); i++ {
		if *mynums[i] > max {
			max = *mynums[i]
		}

		if *mynums[i] < min {
			min = *mynums[i]
		}
	}

	return min, max
}
