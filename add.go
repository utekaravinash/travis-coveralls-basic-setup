package add

// Add returns the sum of numbers passed to it in a slice
func Add(num []int64) (result int64) {

	for _, n := range num {
		result += n
	}

	return result
}
