package chap2_5

func P254(n int) []int {
	var answer []int
	for i := 2; n >= i; i++ {
		if isPrime(i) {
			answer = append(answer, i)
		}
	}
	return answer
}

func isPrime(x int) bool {
	for i := 2; i <= x-1; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
