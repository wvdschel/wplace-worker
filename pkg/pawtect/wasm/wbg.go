package wasm

import (
	"fmt"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func (o *Wasm) Registerwbg(importObj *wasmer.ImportObject) {
	// Register the module
	importObj.Register("wbg", map[string]wasmer.IntoExtern{
		"__wbg_crypto_574e78ad8b13b65f": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_crypto_574e78ad8b13b65f,
		),
		"__wbindgen_is_object": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.I32),
			),
			o.wbg___wbindgen_is_object,
		),
		"__wbg_process_dc0fbacc7c1c06f7": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_process_dc0fbacc7c1c06f7,
		),
		"__wbg_versions_c01dfd4722a88165": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_versions_c01dfd4722a88165,
		),
		"__wbg_node_905d3e251edff8a2": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_node_905d3e251edff8a2,
		),
		"__wbindgen_is_string": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.I32),
			),
			o.wbg___wbindgen_is_string,
		),
		"__wbg_require_60cc747a6bc5215a": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_require_60cc747a6bc5215a,
		),
		"__wbindgen_is_function": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.I32),
			),
			o.wbg___wbindgen_is_function,
		),
		"__wbindgen_string_new": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.I32, wasmer.I32),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbindgen_string_new,
		),
		"__wbg_msCrypto_a61aeb35a24c1329": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_msCrypto_a61aeb35a24c1329,
		),
		"__wbg_randomFillSync_ac0988aba3254290": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef, wasmer.AnyRef),
				wasmer.NewValueTypes(),
			),
			o.wbg___wbg_randomFillSync_ac0988aba3254290,
		),
		"__wbg_getRandomValues_b8f5dbd5f3995a9e": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef, wasmer.AnyRef),
				wasmer.NewValueTypes(),
			),
			o.wbg___wbg_getRandomValues_b8f5dbd5f3995a9e,
		),
		"__wbg_newnoargs_105ed471475aaf50": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.I32, wasmer.I32),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_newnoargs_105ed471475aaf50,
		),
		"__wbg_call_672a4d21634d4a24": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef, wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_call_672a4d21634d4a24,
		),
		"__wbindgen_is_undefined": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.I32),
			),
			o.wbg___wbindgen_is_undefined,
		),
		"__wbg_call_7cccdd69e0791ae2": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef, wasmer.AnyRef, wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_call_7cccdd69e0791ae2,
		),
		"__wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(),
				wasmer.NewValueTypes(wasmer.I32),
			),
			o.wbg___wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0,
		),
		"__wbg_static_accessor_SELF_37c5d418e4bf5819": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(),
				wasmer.NewValueTypes(wasmer.I32),
			),
			o.wbg___wbg_static_accessor_SELF_37c5d418e4bf5819,
		),
		"__wbg_static_accessor_WINDOW_5de37043a91a9c40": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(),
				wasmer.NewValueTypes(wasmer.I32),
			),
			o.wbg___wbg_static_accessor_WINDOW_5de37043a91a9c40,
		),
		"__wbg_static_accessor_GLOBAL_88a902d13a557d07": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(),
				wasmer.NewValueTypes(wasmer.I32),
			),
			o.wbg___wbg_static_accessor_GLOBAL_88a902d13a557d07,
		),
		"__wbg_buffer_609cc3eee51ed158": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_buffer_609cc3eee51ed158,
		),
		"__wbg_newwithbyteoffsetandlength_d97e637ebe145a9a": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef, wasmer.I32, wasmer.I32),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_newwithbyteoffsetandlength_d97e637ebe145a9a,
		),
		"__wbg_new_a12002a7f91c75be": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_new_a12002a7f91c75be,
		),
		"__wbg_set_65595bdd868b3009": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef, wasmer.AnyRef, wasmer.I32),
				wasmer.NewValueTypes(),
			),
			o.wbg___wbg_set_65595bdd868b3009,
		),
		"__wbg_newwithlength_a381634e90c276d4": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.I32),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_newwithlength_a381634e90c276d4,
		),
		"__wbg_subarray_aa9065fa9dc5df96": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef, wasmer.I32, wasmer.I32),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbg_subarray_aa9065fa9dc5df96,
		),
		"__wbindgen_throw": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.I32, wasmer.I32),
				wasmer.NewValueTypes(),
			),
			o.wbg___wbindgen_throw,
		),
		"__wbindgen_memory": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			o.wbg___wbindgen_memory,
		),
		"__wbindgen_init_externref_table": wasmer.NewFunction(
			o.Store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(),
				wasmer.NewValueTypes(),
			),
			o.wbg___wbindgen_init_externref_table,
		),
	})
}

func (o *Wasm) wbg___wbg_crypto_574e78ad8b13b65f(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_crypto_574e78ad8b13b65f(%v)\n", args)
	// TODO implement import '__wbg_crypto_574e78ad8b13b65f'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbindgen_is_object(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbindgen_is_object(%v)\n", args)
	// TODO implement import '__wbindgen_is_object'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_process_dc0fbacc7c1c06f7(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_process_dc0fbacc7c1c06f7(%v)\n", args)
	// TODO implement import '__wbg_process_dc0fbacc7c1c06f7'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_versions_c01dfd4722a88165(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_versions_c01dfd4722a88165(%v)\n", args)
	// TODO implement import '__wbg_versions_c01dfd4722a88165'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_node_905d3e251edff8a2(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_node_905d3e251edff8a2(%v)\n", args)
	// TODO implement import '__wbg_node_905d3e251edff8a2'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbindgen_is_string(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbindgen_is_string(%v)\n", args)
	// TODO implement import '__wbindgen_is_string'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_require_60cc747a6bc5215a(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_require_60cc747a6bc5215a(%v)\n", args)
	// TODO implement import '__wbg_require_60cc747a6bc5215a'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbindgen_is_function(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbindgen_is_function(%v)\n", args)
	// TODO implement import '__wbindgen_is_function'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbindgen_string_new(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbindgen_string_new(%v)\n", args)
	// TODO implement import '__wbindgen_string_new'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_msCrypto_a61aeb35a24c1329(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_msCrypto_a61aeb35a24c1329(%v)\n", args)
	// TODO implement import '__wbg_msCrypto_a61aeb35a24c1329'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_randomFillSync_ac0988aba3254290(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_randomFillSync_ac0988aba3254290(%v)\n", args)
	// TODO implement import '__wbg_randomFillSync_ac0988aba3254290'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_getRandomValues_b8f5dbd5f3995a9e(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_getRandomValues_b8f5dbd5f3995a9e(%v)\n", args)
	// TODO implement import '__wbg_getRandomValues_b8f5dbd5f3995a9e'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_newnoargs_105ed471475aaf50(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_newnoargs_105ed471475aaf50(%v)\n", args)
	// TODO implement import '__wbg_newnoargs_105ed471475aaf50'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_call_672a4d21634d4a24(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_call_672a4d21634d4a24(%v)\n", args)
	// TODO implement import '__wbg_call_672a4d21634d4a24'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbindgen_is_undefined(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbindgen_is_undefined(%v)\n", args)
	// TODO implement import '__wbindgen_is_undefined'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_call_7cccdd69e0791ae2(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_call_7cccdd69e0791ae2(%v)\n", args)
	// TODO implement import '__wbg_call_7cccdd69e0791ae2'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0(%v)\n", args)
	// TODO implement import '__wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_static_accessor_SELF_37c5d418e4bf5819(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_static_accessor_SELF_37c5d418e4bf5819(%v)\n", args)
	// TODO implement import '__wbg_static_accessor_SELF_37c5d418e4bf5819'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_static_accessor_WINDOW_5de37043a91a9c40(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_static_accessor_WINDOW_5de37043a91a9c40(%v)\n", args)
	// TODO implement import '__wbg_static_accessor_WINDOW_5de37043a91a9c40'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_static_accessor_GLOBAL_88a902d13a557d07(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_static_accessor_GLOBAL_88a902d13a557d07(%v)\n", args)
	// TODO implement import '__wbg_static_accessor_GLOBAL_88a902d13a557d07'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_buffer_609cc3eee51ed158(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_buffer_609cc3eee51ed158(%v)\n", args)
	// TODO implement import '__wbg_buffer_609cc3eee51ed158'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_newwithbyteoffsetandlength_d97e637ebe145a9a(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_newwithbyteoffsetandlength_d97e637ebe145a9a(%v)\n", args)
	// TODO implement import '__wbg_newwithbyteoffsetandlength_d97e637ebe145a9a'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_new_a12002a7f91c75be(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_new_a12002a7f91c75be(%v)\n", args)
	// TODO implement import '__wbg_new_a12002a7f91c75be'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_set_65595bdd868b3009(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_set_65595bdd868b3009(%v)\n", args)
	// TODO implement import '__wbg_set_65595bdd868b3009'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_newwithlength_a381634e90c276d4(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_newwithlength_a381634e90c276d4(%v)\n", args)
	// TODO implement import '__wbg_newwithlength_a381634e90c276d4'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbg_subarray_aa9065fa9dc5df96(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbg_subarray_aa9065fa9dc5df96(%v)\n", args)
	// TODO implement import '__wbg_subarray_aa9065fa9dc5df96'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbindgen_throw(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbindgen_throw(%v)\n", args)
	// TODO implement import '__wbindgen_throw'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbindgen_memory(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbindgen_memory(%v)\n", args)
	// TODO implement import '__wbindgen_memory'
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

func (o *Wasm) wbg___wbindgen_init_externref_table(args []wasmer.Value) ([]wasmer.Value, error) {
	fmt.Printf("> wasm import call: wbg.__wbindgen_init_externref_table(%v)\n", args)
	// TODO implement import '__wbindgen_init_externref_table'
	t := o.WbindgenExport2()
	// TODO can't grow table from go?
	_ = t
	return []wasmer.Value{wasmer.NewValue(0.0, wasmer.F64)}, nil
}
