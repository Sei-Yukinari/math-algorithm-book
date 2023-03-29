package chap2_3

import "math"

func P2331(x int) int {
	return 2*x + 3
}

func P2332(x int) int {
	return int(math.Pow(10, float64(x)))
}

func P2333(x int) int {
	return int(math.Log(float64(x)) / math.Log(4))
}

func P2334(x int) int {
	return int(math.Log(math.Sqrt(float64(x))) / math.Log(4))
}
