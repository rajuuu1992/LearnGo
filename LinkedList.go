package ll

import "log"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

// func main() {
// 	head := &LinkedList{1, nil}
// 	head.Add(2)
// 	head.Add(3)
// 	head.Add(4)
// 	head.Add(5)
// 	head.Print()
// 	log.Printf(" Sum = %v", head.Sum())

// }

func (head *LinkedList) Sum() int {
	if head == nil {
		return 0
	}

	return head.Val + head.Next.Sum()
}

func (head *LinkedList) Print() {
	for iter := head; iter != nil; iter = iter.Next {
		log.Printf(" --> %v ", iter.Val)
	}
}

func (head *LinkedList) Add(val int) {

	node := &LinkedList{val, nil}

	if head == nil {
		head = node
		// return head
	}

	iter := head
	for ; iter.Next != nil; iter = iter.Next {

	}

	iter.Next = node
	// return head
}
