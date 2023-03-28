package main

import "fmt"

type C2022Result struct {
	and int
	or  int
	xor int
}

func C2022(a, b int) C2022Result {
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	return C2022Result{
		and: a & b,
		or:  a | b,
		xor: a ^ b,
	}
}
