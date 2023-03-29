package main

type C2031Result struct {
	a1 int
	a2 int
	a3 int
}

func func1() int {
	return 2021
}

var cnt int = 1000

func func2(pos int) int {
	cnt += 1
	return cnt + pos
}

func C2031() C2031Result {
	return C2031Result{
		a1: func1(),
		a2: func2(500),
		a3: func2(500),
	}
}
