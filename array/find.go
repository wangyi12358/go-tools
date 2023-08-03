package array

func Find[T interface{}](array []T, callback func(T) bool) *T {
	for _, v := range array {
		if callback(v) {
			return &v
		}
	}
	return nil
}
