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

	test := []byte{'A', 'B', 'B', 'A', 'A', 'C', 'C'}

	newByte := []byte{}

	temp := byte(0)
	for _, value := range test{
		if value != temp{
			newByte = append(newByte, value)
		}
		temp = value
	}

	fmt.Printf("%c\n", newByte)

	fmt.Println(StringReverser("Google"))

}