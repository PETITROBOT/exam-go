package exam

func Average(array []int) int {
	result := 0
	count := 0
	for _, i := range array {
		if i == -1 {

		} else {
			result += i
			count++
		}
	}
	result = result / count
	return result
}
