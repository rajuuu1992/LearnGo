package main

import "fmt"

func newfn(i int) (a int, b int) {
	defer func() {
		fmt.Println("Exiting...")
	}()
	if i == 0 {
		fmt.Println("Inside newfncd ..%d", i)
		return
	}
	return 1,1
}

//  args are slices
func varFunc(arg ...string) {
	fmt.Println("Entering VarFunc")
	defer func() {
		fmt.Println("Exiting 1...")
	}()
	for _, n := range arg {
		fmt.Print("n = ", n)
	}

	defer func() {
		fmt.Println("Exiting 2...")
	}()
}

func callbackFn(x func(int) int) {
	fmt.Printf("-- CallbackFn....")
	x(666)
}

func Panic(f func()) (b bool) {
	defer func() {
		if x:= recover() ; x !=nil {
			fmt.Println("--Recover called", x)
			b  = true
		}
	}()
	f()
	return
}

func main() {
	a := func(i int) int {
        fmt.Println("Inside newfncd ..%d", i)
		return i
	}

	fmt.Println("a = ", a)
	
	fmt.Println(newfn(1))
	fmt.Println(newfn(0))
	varFunc("a", "b", "c", "d")
	varFunc("aa", "bb", "cc", "dd", "ee")

	callbackFn(a)

	outOfBoundFn := func () {
		var array [100]int
		slice := array[:10]

		fmt.Printf("---Before 100")
		slice[100]  = 10
		fmt.Printf("-----After 100")
	}
	Panic(outOfBoundFn)
}