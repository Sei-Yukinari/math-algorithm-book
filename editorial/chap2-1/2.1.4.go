package chap2_2

import (
	"math"
	"strconv"
)

type P214Result struct {
	decimalNumber int
	nNumber       NNumber
}

type NNumber struct {
	binaryNumber  string
	ternaryNumber string
}

func P214(n1, n2 string) P214Result {
	decimal := binaryToDecimal(n1)
	return P214Result{
		decimalNumber: decimal,
		nNumber: NNumber{
			binaryNumber:  decimalToNNumber(n2, 2),
			ternaryNumber: decimalToNNumber(n2, 3),
		},
	}
}

func binaryToDecimal(n string) int {
	decimal := 0
	for i, v := range n {
		digit, _ := strconv.Atoi(string([]rune{v}))
		decimal += digit * int(math.Pow(2, float64(len(n)-1-i)))
	}
	return decimal
}

func decimalToNNumber(n string, baseNumber int) string {
	decimal, _ := strconv.Atoi(n)
	answer := ""
	for decimal > 0 {
		answer += strconv.Itoa(decimal % baseNumber)
		decimal = decimal / baseNumber
	}
	return reverse(answer)
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
