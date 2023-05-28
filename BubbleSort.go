package main

import "fmt"


func main() {
	arr := [...] int {10, 2, 11, 4, 0}

	fmt.Println("Before  =", arr)
	for i, _ :=  range arr {
		for j := i+1; j < len(arr); j++ {
			if (arr[i] > arr[j]) {
				arr[i], arr[j] =arr[j], arr[i]
			}
		}
	}

	fmt.Println("After  =", arr)
	    
}