package iterable

import (
	"math"
)

func Zip(payload [][]string) [][]string {
	result := make([][]string, 0)
	var min int
	for i := 0; i < len(payload)-1; i++ {
		min = int(math.Min(float64(len(payload[i])), float64(len(payload[i+1]))))
	}
	for i := 0; i < min; i++ {
		tmp := make([]string, 0)
		for j := 0; j < len(payload); j++ {
			tmp = append(tmp, payload[j][i])
		}
		result = append(result, tmp)
	}
	return result
}
