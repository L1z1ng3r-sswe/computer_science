func intToSlice(num int) []int {
	var res []int

	for num > 0 {
		digit := num % 10
		num /= 10
		res = append([]int{digit}, res...)
	}

	return res
}

func sliceToInt(digits []int) int {
	var num int

	for _, digit := range digits {
		num *= 10
		num += digit
	}

	return num
}