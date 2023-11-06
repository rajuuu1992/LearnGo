package main

import (
	"fmt"
	"flag"
	"strings"
)

var n = flag.Bool("n", false, "This is not an option")
var sep = flag.String("S", " ", " ")
var sep2 = flag.String("S2", "2nd string arg", "Hi enter 2nd string arg")

func main() {
    flag.Parse()

	fmt.Println(strings.Join(flag.Args(), ", "))

	
	if !*n {
		fmt.Println(" N  = %v", *n )
	}
	fmt.Println(" S  = " +  *sep )
	fmt.Println(" S2  = " +  *sep2 )
}