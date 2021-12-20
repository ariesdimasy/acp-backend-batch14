package answer

import (
	"math"
)

func MoneyChange(money int) []int {
	var moneys = []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000}
	var result = []int{}

	for i := len(moneys) - 1; i >= 0; i-- {
		if money > moneys[i] {
			var temp float64 = math.Floor(float64(money) / float64(moneys[i]))
			money = money - (int(temp) * moneys[i])

			for j := 0; j < int(temp); j++ {
				result = append(result, moneys[i])
			}
		}
	}

	return result
}
