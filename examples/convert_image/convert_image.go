package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dolmen-go/kittyimg"
	"github.com/jrsap/wplace-worker/pkg/wplace"
)

func main() {
	for _, a := range os.Args[1:] {
		fmt.Println(a)
		img, err := wplace.LoadImage(a)
		if err != nil {
			log.Fatal(err)
		}
		img = wplace.ConvertToPallette(img)

		kittyimg.Fprintln(os.Stdout, img)
	}
}
