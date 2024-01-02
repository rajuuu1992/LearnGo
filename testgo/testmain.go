package main

func Add(a, b int) int {
	return a + b
}

func Add2(a, b int) int {
	res := a

	for i := 0; i < b; i++ {
		res = res + 1
	}
	return res
}

func mul(a, b int) int {
	return a * b
}

func sub(a, b int) int {
	return a - b
}
