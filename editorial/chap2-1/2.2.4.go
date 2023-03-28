package chap2_1

func P2224(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum % 100
}
