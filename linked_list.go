package main

import "container/list"

func LinkedListCreator(n int) *list.List {
	newList := list.New()

	for i := 1; i <= n; i++{
		newList.PushBack(i)
	}

	return newList
}
