package main

func C2013(n []int) int {
	answer := 0
	for _, v := range n {
		answer += v
	}
	return answer
}
