package main

import (
	"fmt"
	"flag"
	"strings"
	"encoding/json"
)

type ListFlag []string

func (f *ListFlag) MarshalText() ([]byte, error) {
	fmt.Println("JSON Marshall %v", f)
	return json.Marshal(*f)
}

func (f *ListFlag) UnmarshalText(b []byte) error {
	for _, str := range strings.Split(string(b), ",") {
		*f = append(*f, str)
	}
	return nil
}

func main() {
	list := ListFlag([]string{"foo", "bar"})
	flag.TextVar(&list, "list", &list, "your flag usage")

	flag.Parse()
	fmt.Println(list)
}