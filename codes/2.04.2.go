package main

func C2042(n, x, y int) int {
	cnt := 0
	for i := 1; i <= n; i++ {
		if i%x == 0 || i%y == 0 {
			cnt += 1
		}
	}
	return cnt
}
