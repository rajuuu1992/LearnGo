package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type ByteCounter int
type WordCounter int

func (counter *ByteCounter) Write(bytes []byte) (int, error) {
	*counter += ByteCounter(len(bytes))

	return len(bytes), nil
}

func (counter *WordCounter) Write(words []string) (int, error) {
	*counter = WordCounter(len(words))

	return len(words), nil
}

type CountingWriter struct {
	writer       io.Writer
	bytesWritten int
}

func (writer *CountingWriter) Write(b []byte) int {
	n, err := writer.writer.Write(b)
	if err != nil {
		return 0
	}
	writer.bytesWritten = n
	return n
}

func (writer *CountingWriter) BytesWritten() int {
	return writer.bytesWritten
}

func NewCountingWriter(w io.Writer) *CountingWriter {
	return &CountingWriter{
		writer:       w,
		bytesWritten: 0,
	}
}

func main() {

	var c ByteCounter
	c.Write([]byte("goprogramminglanguage"))
	fmt.Println(c)

	c = 0
	name := "Newnmae"
	fmt.Fprintf(&c, "hello %s", name)
	fmt.Println(c)

	// input := bufio.NewScanner(os.Stdin)
	// input.Split(bufio.ScanWords)

	// var words []string

	// for input.Scan() {
	// 	words = append(words, input.Text())
	// }

	// var wc WordCounter
	// wc.Write(words)
	// fmt.Println(wc)

	cw := NewCountingWriter(os.Stdout)

	cw.Write([]byte("Hello boss how are you"))
	fmt.Println(cw.BytesWritten())

	cw.Write([]byte("This counts number of bytes written in console"))
	fmt.Println(cw.BytesWritten())

	r := strings.NewReader("this is a sample string reader")

	res := LimitReader(r, 10)

	_, err := io.Copy(os.Stdout, res)

	if err != nil {
		panic(err)
	}

	// fmt.Println("OP =%v", op)

	var w io.Writer

	fmt.Printf(" w= %v, Type of w = %T", w, w)

	w = os.Stdout

	fmt.Printf(" w= %v, Type of w = %T", w, w)

	w = new(bytes.Buffer)
	fmt.Printf(" w= %v, Type of w = %T", w, w)

}

type LimitedReader struct {
	R   io.Reader
	max int64
}

func (lr *LimitedReader) Read(p []byte) (n int, err error) {
	if lr.max <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > lr.max {
		p = p[0:lr.max]
	}

	n, err = lr.R.Read(p)
	lr.max -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
