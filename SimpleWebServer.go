package main

import (
	"net/http"
	"fmt"
	"log"
	"sync"	
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"io"
)

var mu sync.Mutex
var count int


func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	log.Fatal(http.ListenAndServe("localhost:8765", nil))
}
func handler(w http.ResponseWriter, r* http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "Method: %q URL = %q  Proto = %v\n", r.Method, r.URL, r.Proto)
    fmt.Fprintf(w, "Host: %q RemoteAddr = %q\n", r.Host, r.RemoteAddr)
	for k, v := range r.Header {
		fmt.Fprintf(w, " Hdr %v = %v\n", k, v)
	}

    if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form [%q] = %q\n", k, v)
	}

	// gifPattern(w)
}

func counter(w http.ResponseWriter, r* http.Request) {
	
	mu.Lock()
	fmt.Fprintf(w, "Received Req Count = %v\n", count)
	mu.Unlock()
}


var palette = []color.Color{color.White, color.Black}

const (
	wIndex = 0
	bIndex = 1
)

func gifPattern(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0	
	anim := gif.GIF{LoopCount : nframes}
	phase := 0.0

	for i:=0 ; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size +1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t+= res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x*size+0.5), size + int(y*size + 0.5), bIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
		gif.EncodeAll (out, &anim)
	}
}
