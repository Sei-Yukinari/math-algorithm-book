package main

func C2043(n, s int) int {
	answer := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i+j <= s {
				answer += 1
			}
		}
	}
	return answer
}
