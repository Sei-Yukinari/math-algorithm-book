package chap2_3

import "math"

func P2321() int {
	return int(math.Log2(float64(8)))
}

func P2322() int {
	return int(math.Pow(100, 1.5))
}

func P2323() []int {
	var result []int
	result = append(result, int(math.Floor(20.21)))
	result = append(result, int(math.Ceil(20.21)))
	return result
}
