package answer

type Student struct {
	Name  []string
	Score []int
}

func (s Student) Average() float64 {
	var avg float64
	var total float64 = 0

	for i := 0; i < len(s.Score); i++ {
		total = total + float64(s.Score[i])
	}

	avg = total / float64(6)

	return avg
}

func (s Student) Min() (min int, name string) {
	var finalMin int = s.Score[0]
	var finalName string = s.Name[0]

	for i := 1; i < len(s.Name); i++ {
		if finalMin > s.Score[i] {
			finalMin = s.Score[i]
			finalName = s.Name[i]
		}
	}

	return finalMin, finalName
}

func (s Student) Max() (max int, name string) {
	var finalMax int = s.Score[0]
	var finalName string = s.Name[0]

	for i := 1; i < len(s.Name); i++ {
		if finalMax < s.Score[i] {
			finalMax = s.Score[i]
			finalName = s.Name[i]
		}
	}

	return finalMax, finalName
}
