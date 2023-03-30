package main

import "sort"

func C3013(n int) []int {
	var answer []int
	for i := 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		answer = append(answer, i)
		if i != n/i {
			answer = append(answer, n/i)
		}
	}
	sort.Ints(answer)
	return answer
}
