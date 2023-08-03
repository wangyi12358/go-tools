package array

func Map[T interface{}, R interface{}](array []T, callback func(T) R) []R {
	var result []R
	for _, v := range array {
		result = append(result, callback(v))
	}
	return result
}
