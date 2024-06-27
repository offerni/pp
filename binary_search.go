package pp

func IntBinarySearch(term int, data []int) bool {
	start := 0
	end := len(data) - 1

	for start <= end {
		mid := (start + end) / 2
		midVal := data[mid]

		if term == midVal {
			return true
		}

		if term < midVal {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}
