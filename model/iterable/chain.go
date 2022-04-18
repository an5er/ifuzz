package iterable

func Chain(payload [][]string) []string {
	result := make([]string, 0)
	for i := 0; i < len(payload); i++ {
		for j := 0; j < len(payload[i]); j++ {
			result = append(result, payload[i][j])
		}
	}
	return result
}
