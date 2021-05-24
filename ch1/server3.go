package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%s %s %s \n", request.Method, request.URL, request.Proto)

	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(writer, "Host = %q\n", request.Host)
	fmt.Fprintf(writer, "RemoteAddr = %q\n", request.RemoteAddr)

	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range request.Form {
		fmt.Fprintf(writer, "Form[%q] = %q\n", k, v)
	}
}

func lissajous(w http.ResponseWriter, request *http.Request) {
	const (
		cycles   = 5
		res      = 0.001
		size     = 100
		nframews = 64
		delay    = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframews}
	phase := 0.0

	for i := 0; i < nframews; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += .1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
