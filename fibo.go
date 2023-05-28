package main

import "fmt"

func fibo(num int) (res int) {
	if num <= 2 {
		return num - 1
	}

	x:= 0
	y:= 1

	for i := 3; i<=num;i++ {
		res = (x + y)
		x = y
		y = res
		fmt.Printf("----%d, ", res)
	}
	// fmt.Println(x, y)
	return res
}

func main() {
	
	fmt.Println(fibo(50))
}