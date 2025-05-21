// count number of ones:
func countBitOne(n int) int {
	var res int
	for n > 0 {
		if n&0b00000001 == 1 {
			res++
		}
		n = n >> 1
	}

	return res
}

func isSet(n int, i int) bool {
	return (n>>i)&1 == 1
}

func setBit(n int, i int) int {
	return n | (1 << i)
}

func clearBit(n int, i int) int {
	return n &^ (1 << i)
}

func toggleBit(n int, i int) int {
	return n & !(1 << i)
}

func convertBinToDec(num int) int {
	var res int
	var c int

	for num > 0 {
		res += (num & 0b1) * int(math.Pow(2, float64(c)))

		c++
		num >>= 1
	}

	return res
}