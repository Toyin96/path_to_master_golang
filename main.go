package main

import (
	"fmt"
)

func task(a...int) int {
	if len(a) == 0 {
		return 0
	}
	value := a[0]
	fmt.Println(value)
	for _, v := range a {
		if v < value {
			value = v
		}
	}
	return value
}


func main() {

	value := LinkedListCreator(4)
	fmt.Println(*value)
}