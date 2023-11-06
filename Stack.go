package main

import "fmt"
var top = -1

func Push(arr []int, elem int) {
	defer func() {
		fmt.Println("Stack = ", arr)
	}()

	if len(arr) >= cap(arr) {
		fmt.Println("Stack FUll , can't push")
		return
	}
	top++
	arr = append(arr, elem)
}

func Pop(arr[] int) {
	defer func() {
		fmt.Println("Stack = ", arr)
	}()

	if len(arr) == 0 {
		fmt.Println("Stack is empty, can't pop")
		return
	}

	arr = arr[:len(arr)-1]
}
func main() {
	arr:= make([]int, 0, 10)

	fmt.Println("Bef Array = ", arr)
	Pop(arr)
    Push(arr, 10)
	Push(arr, 11)
	Push(arr, 111)
	Pop(arr)
	Push(arr, 1111)
	Pop(arr)

	
	    
}