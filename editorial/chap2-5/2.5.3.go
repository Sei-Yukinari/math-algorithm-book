package chap2_5

func P253(n uint) uint {
	var answer uint = 1
	for i := uint(1); n >= i; i++ {
		answer *= i
	}
	return answer
}
