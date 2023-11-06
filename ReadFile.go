package main

import (
	 "fmt"
	 "io/ioutil"
	 "os"
	 "strings"
	 "bufio"
)


func main() {

	res := make(map[string]int)
	for _, arg := range os.Args[1:] {
		// f, err := os.Open(arg)
		// if err != nil {
		// 	// fmt.Println(" Error = " + err)
		// }
		data, err := ioutil.ReadFile(arg)
		if err != nil {

		}
		for _, str := range strings.Split(string(data), "\n") {
			fmt.Println(" **** str = " + str)
			res[str]++
		}
		// countLines(f, res)
	}
	// strr := "JOINED "
	// for i, val := range os.Args {
	// 	strr += string(i)
	// 	strr += " at " + val
	// 	strr += ", "

	// }

	// lines := make(map[string]int)

	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	lines[input.Text()]++
	// }

	// fmt.Println("\nJOIN + " + strings.Join(os.Args[0:], " "))

	// counts := make(map[string]int)

	// for _, filename := range os.Args[1:] {
	// 	data, err := ioutil.ReadFile(filename)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Dup3 %v", err);
	// 		continue;
	// 	}
	// 	for _, line := range strings.Split(string(data), "\n") {
	// 		counts[line]++
	// 	}
	// }

	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("\n\n '%d'\t : '%s' ..", n, line)
	// 	}
	// }
}

func countLines (f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	
	for input.Scan() {
		fmt.Println("File content = " + input.Text())
		counts[input.Text()]++
	}

}
