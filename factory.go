package pp

// extremely slow
func CreateLargeSliceAppend(size int) []int {
	var slice []int
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	return slice
}

// much quicker
func CreateLargeSliceMake(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	return slice
}
