package main

func StringReverser(s string)string{
	chars := []byte{}

	for i := len(s)-1; i >= 0; i--{
		chars = append(chars, s[i])
	}

	return string(chars)
}
