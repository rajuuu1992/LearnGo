package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	mp := map[string]int{
		"Helo": 1,
		"boss": 2,
	}

	fmt.Println(mp)

	var keys []string
	for key := range mp {
		fmt.Printf(" Key = %v", key)
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("  %v, ", mp[k])
	}

	fmt.Println(keys)

	if num, ok := mp["Helo"]; !ok {
		fmt.Println("hello not there")
	} else {
		fmt.Println("Hello entry present %v", num)
	}

	mp2 := map[string]int{
		"Helo": 1,
		"boss": 2,
	}

	fmt.Println("map equal ", mapEqual(mp, mp2))

	wordFreq("hello.txt")
}

func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, val1 := range x {
		if val2, ok := y[k]; !ok || val1 != val2 {
			return false
		}
	}
	return true
}

func wordFreq(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(" Error = ", err)
		return
	}

	words := make(map[string]int)
	fmt.Println(" IS words map nil = ", (words == nil))
	fmt.Println(" Len words map  = ", len(words))

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		fmt.Printf("Type = %T, val = %v", word, word)

		words[word]++
	}

	for k, v := range words {
		fmt.Println(" %v, %v ", k, v)
	}
}
