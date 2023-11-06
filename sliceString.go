package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0

	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	arr_str := []string{"hello", " alskdlsdkfj", "", "alskdfjlskjdf", "lksjflksdjf"}

	fmt.Println(nonempty(arr_str))

	stk := make([]int, 1)
	stk = append(stk, 2)
	stk = append(stk, 4)

	fmt.Println("Stack = ", stk)

	stk = stk[:len(stk)-1]

	fmt.Println("---------- Stack = ", stk)

	stk = remSlice(stk, 0)

	fmt.Println("---------- RemSlice Stack = ", stk)

	var x []int

	for i := 0; i < 10; i++ {
		x = append(x, i)
		fmt.Println(" %v  Cap = %v, len = %v", x, len(x), cap(x))
	}

	x = append(x, 2, 3, 4)
	fmt.Println("Orig = ", x)
	reverse(x)
	fmt.Println("Reverse = ", x)

	x = appendInt(x, 444)
	fmt.Println("AppendInt = ", x)

	y := make([]int, 2)
	y = appendInt(y, 333)
	fmt.Println("AppendInt333 Y = %v, Cap = %v", y, cap(y))
	y = appendInt(y, 444)
	fmt.Println("AppendInt444 Y = %v, Cap = %v", y, cap(y))
	y = appendInt(y, 555)

	fmt.Println("AppendInt555 Y = %v, Cap = %v", y, cap(y))

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendInt(s []int, val int) []int {

	var z []int
	if len(s)+1 <= cap(s) {
		z = s[:len(s)+1]
	} else {
		zcap := len(s) + 1
		if len(s)+1 <= 2*cap(s) {
			zcap = 2 * len(s)
		}
		z = make([]int, len(s)+1, zcap)
		copy(z, s)
	}
	z[len(s)] = val
	return z
}

func remSlice(slice []int, i int) []int {
	copy(slice[:i], slice[i+1:])
	return slice[:len(slice)-1]
}