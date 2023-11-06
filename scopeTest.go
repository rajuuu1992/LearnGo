package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const (
	usage = `there can be many lnies in this string
	            , this is called raw string, no 
				 escape sequences will be processed here`

	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

func main() {
	x := "strsss!"

	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x = x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}

	// f, err := os.Open("file")
	// if err != nil {
	// 	return
	// }

	// f.Stat()
	// f.Close()

	// var err error
	// var cwd string

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed")
	}
	fmt.Println(" Cwd = %v", cwd)
	// fmt.Println(" Basename = %v", os.basename(cwd))

	str := "hello Boss"

	sub1 := str[0:5]
	sub2 := str[5:]

	rr := []rune(sub2)
	rr[0] = 'z'
	sub2 = string(rr)

	fmt.Println(" Str = %s, sub1 = %s, sub2 = %s", str, sub1, sub2)

	fmt.Println(" Raw string = ", usage)

	s := "hello"
	prefix := "hell"
	fmt.Println(" IsPrefix = %s, %s  = %v", s, prefix, IsPrefix(s, prefix))

	s = "hihowareyou"
	prefix = "notwell"
	fmt.Println(" IsPrefix = %s, %s  = %v", s, prefix, IsPrefix(s, prefix))

	num := 123456712123123
	fmt.Println(" Num = %v, Comma = %s", num, comma(strconv.Itoa(num)))
	fmt.Println(" Num = %v,         Comma2 = %s", num, comma2(strconv.Itoa(num)))

	fmt.Println(" Anagram = ", IsAnagram("hellohai", "iahlloeh"))

	fmt.Println(" B = %v, KB = %v, GB = %v", B, KB, GB)
}

func IsPrefix(s string, prefix string) bool {
	return len(s) >= len(prefix) &&
		s[:len(prefix)] == prefix
}

func comma(num string) string {

	if len(num) <= 3 {
		return num
	}

	return comma(num[:len(num)-3]) + "," + num[len(num)-3:]
}

func comma2(num string) string {

	if len(num) < 3 {
		return num
	}

	byt := new(bytes.Buffer)

	rel := len(num) % 3
	fmt.Println(" len = %v, rel = %v", len(num), rel)
	count := 0
	for i := 0; i < len(num); i++ {
		if i < rel {
			byt.WriteByte(num[i])
			continue
		}

		if i == rel {
			if rel != 0 {
				byt.WriteByte(',')
			}
			byt.WriteByte(num[i])
			continue
		}

		if i > rel {
			count++
			if count%3 == 0 {
				byt.WriteByte(',')
			}
			byt.WriteByte(num[i])
		}
	}
	fmt.Println(" Res Byt = ", byt)
	return byt.String()

}
func Sort(s string) string {
	chars := []rune(s)

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func IsAnagram(s1 string, s2 string) bool {

	return len(s1) == len(s2) && Sort(s1) == Sort(s2)
}
