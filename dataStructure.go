package main

import (
	"fmt"
	"container/list"
)

func main() {
	var listOfInteger list.List
	listOfInteger.PushBack(10)
	listOfInteger.PushBack(20)
	listOfInteger.PushBack(30)

	for element := listOfInteger.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}