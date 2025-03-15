package utils

func InArray[T comparable](item T, array []T) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}

func ArrayDiff[T comparable](array1 []T, array2 []T) []T {
	var diff []T
	for _, v := range array1 {
		if !InArray(v, array2) {
			diff = append(diff, v)
		}
	}
	return diff
}
