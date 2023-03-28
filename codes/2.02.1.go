package main

import (
	"fmt"
	"math"
)

func main() {
	//四則演算
	fmt.Printf("%d\n", 869+120)
	fmt.Printf("%d\n", 869-120)
	fmt.Printf("%d\n", 869*120)
	fmt.Printf("%d\n", 869/120)

	//余剰
	fmt.Printf("%d\n", 8%5)
	fmt.Printf("%d\n", 869%120)

	//絶対値
	fmt.Printf("%.10g\n", math.Abs(-45))
	fmt.Printf("%.10g\n", math.Abs(15))

	//累乗
	fmt.Printf("%d\n", int(math.Pow(10, 2)))
	fmt.Printf("%d\n", int(math.Pow(3, 4)))

	//ルート
	fmt.Printf("%.5f\n", math.Sqrt(4))
	fmt.Printf("%.5f\n", math.Sqrt(2))
}
