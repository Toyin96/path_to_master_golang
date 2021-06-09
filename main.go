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

	 insertion := []string{"I", "am", "happy"}
	 slice := []string{"Toyin", "is", "age", "school"}

	result := func (slice, insertion []string, index int) []string {
		newSlice := slice[:index]
		fmt.Println(slice)
		copy(slice, insertion)
		fmt.Println(slice)
		after := slice[index:]
		newSlice = append(newSlice, slice...)
		newSlice = append(newSlice, after...)
		return newSlice
	}(slice, insertion, 1)

	fmt.Println(result)

}