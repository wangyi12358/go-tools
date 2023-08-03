package array

func Filter[T interface{}](array []T, callback func(T) bool) []T {
	var result []T
	for _, v := range array {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}
