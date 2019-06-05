package raindrops

import (
	"fmt"
)

func Convert(number int) string {
	factors := []int{3,5,7}
	var result string
	factorMap := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	for _, v := range factors {
		if number % v == 0 {
			result = fmt.Sprintf("%s%s", result, factorMap[v])
		}
	}
	if result == "" {
		result = fmt.Sprintf("%d", number)
	}
	return result
}

