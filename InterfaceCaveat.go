package main

import (
	"bytes"
	"io"
	"log"
)

const debug = true

func main() {
	var buf io.Writer

	if debug {
		log.Printf("Start of Debug")
		buf = new(bytes.Buffer)
	}

	f(buf)

	if debug {
		log.Printf("End of Debug")
	}
}

func f(writer io.Writer) {
	if writer != nil {
		writer.Write([]byte("Done\n"))
	}
}
