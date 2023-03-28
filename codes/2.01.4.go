package main

const baseNNumbers = 2

func C2014(n int) string {
	answer := ""
	for n > 0 {
		if n%baseNNumbers == 0 {
			answer += "0"
		} else if n%baseNNumbers == 1 {
			answer += "1"
		}
		n = n / baseNNumbers
	}
	return answer
}
