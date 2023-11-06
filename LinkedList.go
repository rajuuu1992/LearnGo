package main

import "log"

type LinkedList struct {
	val  int
	next *LinkedList
}

func main() {
	head := &LinkedList{1, nil}
	head.Add(2)
	head.Add(3)
	head.Add(4)
	head.Add(5)
	head.Print()
	log.Printf(" Sum = %v", head.Sum())

}

func (head *LinkedList) Sum() int {
	if head == nil {
		return 0
	}

	return head.val + head.next.Sum()
}

func (head *LinkedList) Print() {
	for iter := head; iter != nil; iter = iter.next {
		log.Printf(" --> %v ", iter.val)
	}
}

func (head *LinkedList) Add(val int) {

	node := &LinkedList{val, nil}

	if head == nil {
		head = node
		// return head
	}

	iter := head
	for ; iter.next != nil; iter = iter.next {

	}

	iter.next = node
	// return head
}
