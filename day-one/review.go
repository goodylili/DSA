package main

import (
	"container/list"
	"fmt"
)

func main() {
	var listOfIntegers list.List
	listOfIntegers.PushBack(10)
	listOfIntegers.PushBack(20)
	listOfIntegers.PushBack(30)

	for elements := listOfIntegers.Front(); elements != nil; elements = elements.Next() {
		fmt.Println(elements.Value.(int))
	}
}
