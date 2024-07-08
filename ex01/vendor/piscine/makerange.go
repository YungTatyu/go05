package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	var slice []int = make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i + min
	}
	return slice
}
