package main

import "fmt"

type Node struct {
	data any
	next *Node
}

type LinkedList struct {
	headNode *Node
	length   int
}

func (list *LinkedList) addElement(element any) {
	newNode := Node{
		data: element,
	}
	if list.length == 0 {
		list.headNode = &newNode
		list.length++
		return
	} else {
		currentNode := list.headNode
		for position := 0; position < list.length; position++ {
			if currentNode.next == nil {
				currentNode.next = &newNode
				list.length++
				return
			}
			currentNode = currentNode.next
		}
	}
	fmt.Println("added")
}

func (list *LinkedList) getElementIndex(element any) {

	currentPosition := list.headNode
	for index := 0; index < list.length; index++ {
		if currentPosition.data == element {
			fmt.Println(index, currentPosition.data)
		}
		currentPosition = currentPosition.next

	}
}

func (list *LinkedList) replaceElement(element, replacer any) any {

	currentPosition := list.headNode
	for index := 0; index < list.length; index++ {
		if currentPosition.data == element {
			currentPosition.data = replacer
			return replacer
		}
		currentPosition = currentPosition.next
	}
	return element
}

func (list *LinkedList) DeleteElement(element any) {
	currentPosition := list.headNode
	if list.length == 0 {
		return
	}
	for index := 0; index < list.length; index++ {
		if currentPosition.data == element {
			currentPosition = nil // set the position to nil
			for currentPosition.next != nil {
				currentPosition = currentPosition.next
			}
		}

	}

	list.length -= 1

}

func (list *LinkedList) PrintList() {
	if list.length == 0 {
		fmt.Print(nil)
	}
	currentPosition := list.headNode
	for index := 0; index < list.length; index++ {
		fmt.Println(currentPosition.data)
		currentPosition = currentPosition.next
	}
}

func main() {
	exampleList := LinkedList{
		headNode: nil,
		length:   0,
	}
	exampleList.addElement(4)
	exampleList.addElement("Golang")
	exampleList.addElement(false)
	sampleArray := [4]int{5, 6, 3, 7}
	exampleList.addElement(sampleArray)
	exampleList.DeleteElement(false)
	exampleList.replaceElement(4, "200")
	exampleList.getElementIndex("200")
	exampleList.DeleteElement(false)
	exampleList.PrintList()
}
