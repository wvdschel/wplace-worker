package main

import (
	"flag"
	"os"

	"github.com/wvdschel/wplace-worker/pkg/wasmergen"
)

var (
	wasmFile = flag.String("wasmfile", "", "wasm file to generate imports for")
	outDir   = flag.String("outdir", "wasm/generated", "output file to write imports to")
)

func main() {
	flag.Parse()

	if *wasmFile == "" {
		panic("wasmfile is required")
	}

	if *outDir == "" {
		panic("outdir is required")
	}

	wasmBytes, err := os.ReadFile(*wasmFile)
	if err != nil {
		panic(err)
	}

	g := wasmergen.New()

	if err := g.GenerateImportStubs(wasmBytes, *outDir); err != nil {
		panic(err)
	}
}
