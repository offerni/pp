package pp

func IntSerialSearch(key int, data []int) bool {
	for _, num := range data {
		if num == key {
			return true
		}
	}
	return false
}
