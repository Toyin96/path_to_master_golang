package main

func Filter(s []int, fn func(int) bool) []int {
	var newSlice []int
	for _, value := range s{
		if fn(value){
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
