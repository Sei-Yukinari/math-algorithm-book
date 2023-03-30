package chap3_1

func P312(n int) []int {
	var answer []int
	i := 2
	for n > 1 {
		if n%i == 0 {
			n /= i
			answer = append(answer, i)
		} else {
			i++
		}
	}
	return answer
}
