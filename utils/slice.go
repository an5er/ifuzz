package utils

func SliceDelete(slice []string, i int) []string {
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}

func SliceAdd(slice []string, str []string) []string {
	slice = append(slice, str...)
	return slice
}
