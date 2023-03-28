package chap2_1

type P2231Result struct {
	answer1 int
	answer2 int
	answer3 int
}

func P22231(x, y int) P2231Result {
	return P2231Result{
		answer1: x & y,
		answer2: x | y,
		answer3: x ^ y,
	}
}

func P22232(a, b, c, d int) int {
	return a | b | c | d
}
