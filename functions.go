package main

type filter func(n int) bool

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func isOdd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true
}

func collection(slice []int, checker filter) (even []int) {
	for _, value := range slice {
		if checker(value) {
			even = append(even, value)
		}
	}
	return
}

