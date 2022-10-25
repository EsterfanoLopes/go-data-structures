package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Println()
}

func (l linkedList) print() {
	var print *node

	for {
		if print == nil && l.head != nil {
			print = l.head
		} else {
			if print == nil || print.next == nil {
				return
			}
			print = print.next
		}
		fmt.Println(print.data)

	}
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	// to delete first item
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		previousToDelete = previousToDelete.next
		// To when there's no value on the list
		if previousToDelete.next == nil {
			return
		}
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 11}
	node5 := &node{data: 7}
	node6 := &node{data: 2}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	// delete an item
	myList.printListData()
	myList.deleteWithValue(16)
	myList.printListData()

	// delete the first item
	myList.deleteWithValue(2)
	myList.printListData()

	// try to delete an item that not exists on the list
	myList.deleteWithValue(100)
	myList.printListData()

	// try to delete from an empty list
	myEmptyList := linkedList{}
	myEmptyList.deleteWithValue(1)
	myEmptyList.printListData()
}
