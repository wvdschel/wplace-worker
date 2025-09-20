package pawtect

import (
	"context"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type wbg struct {
	mod api.Module
}

// wbgModule provides all the imports required by the WebAssembly module
func wbgModule(r wazero.Runtime) api.Module {
	w := &wbg{}
	builder := r.NewHostModuleBuilder("wbg")

	// __wbg_crypto_574e78ad8b13b65f (param externref) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_crypto_574e78ad8b13b65f).
		Export("__wbg_crypto_574e78ad8b13b65f")

	// __wbindgen_is_object (param externref) (result i32)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_is_object).
		Export("__wbindgen_is_object")

	// __wbg_process_dc0fbacc7c1c06f7 (param externref) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_process_dc0fbacc7c1c06f7).
		Export("__wbg_process_dc0fbacc7c1c06f7")

	// __wbg_versions_c01dfd4722a88165 (param externref) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_versions_c01dfd4722a88165).
		Export("__wbg_versions_c01dfd4722a88165")

	// __wbg_node_905d3e251edff8a2 (param externref) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_node_905d3e251edff8a2).
		Export("__wbg_node_905d3e251edff8a2")

	// __wbindgen_is_string (param externref) (result i32)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_is_string).
		Export("__wbindgen_is_string")

	// __wbg_require_60cc747a6bc5215a (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_require_60cc747a6bc5215a).
		Export("__wbg_require_60cc747a6bc5215a")

	// __wbindgen_is_function (param externref) (result i32)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_is_function).
		Export("__wbindgen_is_function")

	// __wbindgen_string_new (param i32 i32) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_string_new).
		Export("__wbindgen_string_new")

	// __wbg_msCrypto_a61aeb35a24c1329 (param externref) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_msCrypto_a61aeb35a24c1329).
		Export("__wbg_msCrypto_a61aeb35a24c1329")

	// __wbg_randomFillSync_ac0988aba3254290 (param externref externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_randomFillSync_ac0988aba3254290).
		Export("__wbg_randomFillSync_ac0988aba3254290")

	// __wbg_getRandomValues_b8f5dbd5f3995a9e (param externref externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_getRandomValues_b8f5dbd5f3995a9e).
		Export("__wbg_getRandomValues_b8f5dbd5f3995a9e")

	// __wbg_newnoargs_105ed471475aaf50 (param i32 i32) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_newnoargs_105ed471475aaf50).
		Export("__wbg_newnoargs_105ed471475aaf50")

	// __wbg_call_672a4d21634d4a24 (param externref externref) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_call_672a4d21634d4a24).
		Export("__wbg_call_672a4d21634d4a24")

	// __wbindgen_is_undefined (param externref) (result i32)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_is_undefined).
		Export("__wbindgen_is_undefined")

	// __wbg_call_7cccdd69e0791ae2 (param externref externref externref) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_call_7cccdd69e0791ae2).
		Export("__wbg_call_7cccdd69e0791ae2")

	// __wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0 (result i32)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0).
		Export("__wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0")

	// __wbg_static_accessor_SELF_37c5d418e4bf5819 (result i32)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_static_accessor_SELF_37c5d418e4bf5819).
		Export("__wbg_static_accessor_SELF_37c5d418e4bf5819")

	// __wbg_static_accessor_WINDOW_5de37043a91a9c40 (result i32)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_static_accessor_WINDOW_5de37043a91a9c40).
		Export("__wbg_static_accessor_WINDOW_5de37043a91a9c40")

	// __wbg_static_accessor_GLOBAL_88a902d13a557d07 (result i32)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_static_accessor_GLOBAL_88a902d13a557d07).
		Export("__wbg_static_accessor_GLOBAL_88a902d13a557d07")

	// __wbg_buffer_609cc3eee51ed158 (param externref) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_buffer_609cc3eee51ed158).
		Export("__wbg_buffer_609cc3eee51ed158")

	// __wbg_newwithbyteoffsetandlength_d97e637ebe145a9a (param externref i32 i32) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_newwithbyteoffsetandlength_d97e637ebe145a9a).
		Export("__wbg_newwithbyteoffsetandlength_d97e637ebe145a9a")

	// __wbg_new_a12002a7f91c75be (param externref) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_new_a12002a7f91c75be).
		Export("__wbg_new_a12002a7f91c75be")

	// __wbg_set_65595bdd868b3009 (param externref externref i32)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_set_65595bdd868b3009).
		Export("__wbg_set_65595bdd868b3009")

	// __wbg_newwithlength_a381634e90c276d4 (param i32) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_newwithlength_a381634e90c276d4).
		Export("__wbg_newwithlength_a381634e90c276d4")

	// __wbg_subarray_aa9065fa9dc5df96 (param externref i32 i32) (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbg_subarray_aa9065fa9dc5df96).
		Export("__wbg_subarray_aa9065fa9dc5df96")

	// __wbindgen_throw (param i32 i32)
	builder.NewFunctionBuilder().
		WithFunc(w.wbindgen_throw).
		Export("__wbindgen_throw")

	// __wbindgen_memory (result externref)
	builder.NewFunctionBuilder().
		WithFunc(w.wbindgen_memory).
		Export("__wbindgen_memory")

	// __wbindgen_init_externref_table ()
	builder.NewFunctionBuilder().
		WithFunc(w.wbindgen_init_externref_table).
		Export("__wbindgen_init_externref_table")

	// Instantiate the module
	module, err := builder.Instantiate(context.Background())
	if err != nil {
		panic(err)
	}

	w.mod = module

	return module
}

func (w *wbg) wbg_crypto_574e78ad8b13b65f(ctx context.Context, ref uintptr) uintptr {
	log.Printf("wbg.__wbg_crypto_574e78ad8b13b65f called")
	// TODO: Implement crypto functionality
	return 0
}

func (w *wbg) wbg_is_object(ctx context.Context, ref uintptr) int32 {
	log.Printf("wbg.__wbindgen_is_object called")
	// TODO: Implement object check
	return 0
}

func (w *wbg) wbg_process_dc0fbacc7c1c06f7(ctx context.Context, ref uintptr) uintptr {
	log.Printf("wbg.__wbg_process_dc0fbacc7c1c06f7 called")
	// TODO: Implement process functionality
	return 0
}

func (w *wbg) wbg_versions_c01dfd4722a88165(ctx context.Context, ref uintptr) uintptr {
	log.Printf("wbg.__wbg_versions_c01dfd4722a88165 called")
	// TODO: Implement versions functionality
	return 0
}

func (w *wbg) wbg_node_905d3e251edff8a2(ctx context.Context, ref uintptr) uintptr {
	log.Printf("wbg.__wbg_node_905d3e251edff8a2 called")
	// TODO: Implement node functionality
	return 0
}

func (w *wbg) wbg_is_string(ctx context.Context, ref uintptr) int32 {
	log.Printf("wbg.__wbindgen_is_string called")
	// TODO: Implement string check
	return 0
}

func (w *wbg) wbg_require_60cc747a6bc5215a(ctx context.Context) uintptr {
	log.Printf("wbg.__wbg_require_60cc747a6bc5215a called")
	// TODO: Implement require functionality
	return 0
}

func (w *wbg) wbg_is_function(ctx context.Context, ref uintptr) int32 {
	log.Printf("wbg.__wbindgen_is_function called")
	// TODO: Implement function check
	return 0
}

func (w *wbg) wbg_string_new(ctx context.Context, ptr uint32, len uint32) uintptr {
	log.Printf("wbg.__wbindgen_string_new called")
	// TODO: Implement string creation
	return 0
}

func (w *wbg) wbg_msCrypto_a61aeb35a24c1329(ctx context.Context, ref uintptr) uintptr {
	log.Printf("wbg.__wbg_msCrypto_a61aeb35a24c1329 called")
	// TODO: Implement msCrypto functionality
	return 0
}

func (w *wbg) wbg_randomFillSync_ac0988aba3254290(ctx context.Context, ref1 uintptr, ref2 uintptr) {
	log.Printf("wbg.__wbg_randomFillSync_ac0988aba3254290 called")
	// TODO: Implement randomFillSync functionality
}

func (w *wbg) wbg_getRandomValues_b8f5dbd5f3995a9e(ctx context.Context, ref1 uintptr, ref2 uintptr) {
	log.Printf("wbg.__wbg_getRandomValues_b8f5dbd5f3995a9e called")
	// TODO: Implement getRandomValues functionality
}

func (w *wbg) wbg_newnoargs_105ed471475aaf50(ctx context.Context, ptr uint32, len uint32) uintptr {
	log.Printf("wbg.__wbg_newnoargs_105ed471475aaf50 called")
	// TODO: Implement newnoargs functionality
	return 0
}

func (w *wbg) wbg_call_672a4d21634d4a24(ctx context.Context, ref1 uintptr, ref2 uintptr) uintptr {
	log.Printf("wbg.__wbg_call_672a4d21634d4a24 called")
	// TODO: Implement call functionality
	return 0
}

func (w *wbg) wbg_is_undefined(ctx context.Context, ref uintptr) int32 {
	log.Printf("wbg.__wbindgen_is_undefined called")
	// TODO: Implement undefined check
	return 0
}

func (w *wbg) wbg_call_7cccdd69e0791ae2(ctx context.Context, ref1 uintptr, ref2 uintptr, ref3 uintptr) uintptr {
	log.Printf("wbg.__wbg_call_7cccdd69e0791ae2 called")
	// TODO: Implement call functionality
	return 0
}

func (w *wbg) wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0(ctx context.Context) int32 {
	log.Printf("wbg.__wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0 called")
	// TODO: Implement GLOBAL_THIS accessor
	return 0
}

func (w *wbg) wbg_static_accessor_SELF_37c5d418e4bf5819(ctx context.Context) int32 {
	log.Printf("wbg.__wbg_static_accessor_SELF_37c5d418e4bf5819 called")
	// TODO: Implement SELF accessor
	return 0
}

func (w *wbg) wbg_static_accessor_WINDOW_5de37043a91a9c40(ctx context.Context) int32 {
	log.Printf("wbg.__wbg_static_accessor_WINDOW_5de37043a91a9c40 called")
	// TODO: Implement WINDOW accessor
	return 0
}

func (w *wbg) wbg_static_accessor_GLOBAL_88a902d13a557d07(ctx context.Context) int32 {
	log.Printf("wbg.__wbg_static_accessor_GLOBAL_88a902d13a557d07 called")
	// TODO: Implement GLOBAL accessor
	return 0
}

func (w *wbg) wbg_buffer_609cc3eee51ed158(ctx context.Context, ref uintptr) uintptr {
	log.Printf("wbg.__wbg_buffer_609cc3eee51ed158 called")
	// TODO: Implement buffer functionality
	return 0
}

func (w *wbg) wbg_newwithbyteoffsetandlength_d97e637ebe145a9a(ctx context.Context, ref uintptr, offset uint32, length uint32) uintptr {
	log.Printf("wbg.__wbg_newwithbyteoffsetandlength_d97e637ebe145a9a called")
	// TODO: Implement newwithbyteoffsetandlength functionality
	return 0
}

func (w *wbg) wbg_new_a12002a7f91c75be(ctx context.Context, ref uintptr) uintptr {
	log.Printf("wbg.__wbg_new_a12002a7f91c75be called")
	// TODO: Implement new functionality
	return 0
}

func (w *wbg) wbg_set_65595bdd868b3009(ctx context.Context, ref1 uintptr, ref2 uintptr, val int32) {
	log.Printf("wbg.__wbg_set_65595bdd868b3009 called")
	// TODO: Implement set functionality
}

func (w *wbg) wbg_newwithlength_a381634e90c276d4(ctx context.Context, length uint32) uintptr {
	log.Printf("wbg.__wbg_newwithlength_a381634e90c276d4 called")
	// TODO: Implement newwithlength functionality
	return 0
}

func (w *wbg) wbg_subarray_aa9065fa9dc5df96(ctx context.Context, ref uintptr, start uint32, end uint32) uintptr {
	log.Printf("wbg.__wbg_subarray_aa9065fa9dc5df96 called")
	// TODO: Implement subarray functionality
	return 0
}

func (w *wbg) wbindgen_throw(ctx context.Context, ptr uint32, len uint32) {
	log.Printf("wbg.__wbindgen_throw called")
	// TODO: Implement throw functionality
}

func (w *wbg) wbindgen_memory(ctx context.Context) uintptr {
	log.Printf("wbg.__wbindgen_memory called")
	// TODO: Implement memory functionality
	return 0
}

func (w *wbg) wbindgen_init_externref_table(ctx context.Context) {
	log.Printf("wbg.__wbindgen_init_externref_table called")
	// TODO: Implement init_externref_table functionality

}
