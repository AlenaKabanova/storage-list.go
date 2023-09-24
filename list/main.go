package main

import (
	list "list/storage"
)

func main() {
	l := list.NewList()
	l.Add(17)
	l.Add(13)
	l.Add(123)
	l.Add(334324)
	l.RemoveByIndex(3)
	l.RemoveByValue(17)
	l.Print()
}
