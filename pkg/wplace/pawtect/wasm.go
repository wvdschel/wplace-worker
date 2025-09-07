package pawtect

import (
	"context"
	_ "embed"
	"log"

	"github.com/tetratelabs/wazero"
)

//go:embed pawtect.wasm
var pawtectWasmModule []byte

func LoadModule() {
	// Load the WebAssembly module
	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	wbgMod := wbgModule(r)
	_ = wbgMod // The module is available for import by the Wasm module

	// Instantiate the guest Wasm into the same runtime. It exports the `add`
	// function, implemented in WebAssembly.
	mod, err := r.InstantiateWithConfig(ctx, pawtectWasmModule, wazero.NewModuleConfig().WithStartFunctions("__wbindgen_start"))
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}

	mod.ExportedFunction("set_user_id").Call(ctx, 12345)
}
