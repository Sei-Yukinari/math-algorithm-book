package main

func C3011(n uint) bool {
	for i := uint(2); i <= n-1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
