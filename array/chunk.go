package array

func Chunk[T interface{}](array []T, size int) [][]T {
	var result [][]T
	for i := 0; i < len(array); i += size {
		end := i + size
		if end > len(array) {
			end = len(array)
		}
		result = append(result, array[i:end])
	}
	return result
}
