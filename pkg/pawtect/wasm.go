package pawtect

import (
	"context"
	_ "embed"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed pawtect.wasm
var pawtectWasmModule []byte

func Load() {
	// Load the WebAssembly module
	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	wasi_snapshot_preview1.MustInstantiate(ctx, r)
	wbgMod := wbgModule(r)
	_ = wbgMod // The module is available for import by the Wasm module

	mod, err := r.InstantiateWithConfig(ctx, pawtectWasmModule, wazero.NewModuleConfig())
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}

	_, err = mod.ExportedFunction("__wbindgen_start").Call(ctx)
	if err != nil {
		log.Panicf("failed to call __wbindgen_start: %v", err)
	}

	_, err = mod.ExportedFunction("set_user_id").Call(ctx, 12345)
	if err != nil {
		log.Panicf("failed to call set_user_id: %v", err)
	}
	res, err := mod.ExportedFunction("get_load_payload").Call(ctx)
	if err != nil {
		log.Panicf("failed to call get_load_payload: %v", err)
	}
	log.Printf("response: %v", res)
}
