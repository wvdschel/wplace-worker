package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/dolmen-go/kittyimg"
	"github.com/jrsap/wplace-worker/pkg/wplace"
)

func main() {
	for i, a := range os.Args[1:] {
		fmt.Println(a)
		img, err := wplace.LoadImage(a)
		if err != nil {
			log.Fatal(err)
		}
		img = wplace.ConvertToPallette(img)

		f, err := os.OpenFile(fmt.Sprintf("converted_%d.png", i), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if err := png.Encode(f, img); err != nil {
			log.Fatal(err)
		}

		kittyimg.Fprintln(os.Stdout, img)
	}
}
