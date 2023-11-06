package main

import "fmt"

func fibo(num int) (res int) {
	if num <= 2 {
		return num - 1
	}

	x:= 0
	y:= 1

	for i := 3; i<=num;i++ {
		x, y = y, x+y
		// res = (x + y)
		// x = y
		// y = res
		fmt.Printf("----%d %d, ", x, y)
	}
	// fmt.Println(x, y)
	return res
}

func gcd (x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	
	// fmt.Println(fibo(50))

	fmt.Println(gcd(100,30))
}