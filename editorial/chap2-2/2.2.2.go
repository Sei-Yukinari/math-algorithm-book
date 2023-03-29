package chap2_2

import "math"

type P2221Result struct {
	answer1 float64
	answer2 float64
}

func P2221(n1, n2 int) P2221Result {
	return P2221Result{
		answer1: math.Sqrt(float64(n1)),
		answer2: math.Pow(float64(n2), 2),
	}
}

type P2222Result struct {
	answer1 int
	answer2 float64
}

func P2222(x1, y1 float64, x2, y2 int) P2222Result {

	rootN := 1 / x1

	return P2222Result{
		answer1: int(math.Pow(y1, rootN)),
		answer2: math.Pow(float64(x2), float64(y2)),
	}
}
