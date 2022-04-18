package iterable

import (
	"log"
)

func Product(sep [][]string) [][]string {
	if len(sep) > 2 {
		log.Fatalf("Product arg only 2,too many")
	}

	combinations := make([][]string, 0)
	countIndex := len(sep) - 1
	counter := make([]int, countIndex+1)
	firstSize := len(sep[0]) - 1
	for firstSize > counter[0] {
		combination := make([]string, 0)
		for i := countIndex; i >= 0; i-- {
			combination = append(combination, sep[i][counter[i]])
		}
		combinations = append(combinations, combination)
		counter[countIndex]++
		for i := countIndex; i >= 0; i-- {
			if counter[i] > len(sep[i])-1 {
				counter[i] = 0
				if i > 0 {
					counter[i-1]++
				}
			}
		}
	}
	return combinations
}
