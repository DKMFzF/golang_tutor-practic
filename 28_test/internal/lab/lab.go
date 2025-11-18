package lab

func Sum(v ...int) int {
	var sum int
	for _, v := range v {
		sum += v
	}

	return sum
}
