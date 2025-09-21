package pawtect

import (
	"log"

	"github.com/wasmerio/wasmer-go/wasmer"
)

// WBG import methods for WebAssembly bindings

// __wbg_crypto_574e78ad8b13b65f - (param externref) (result externref)
func (p *Pawtect) __wbg_crypto_574e78ad8b13b65f(args []wasmer.Value) ([]wasmer.Value, error) {
	log.Printf("__wbg_crypto_574e78ad8b13b65f({})", args)
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

// __wbindgen_is_object - (param externref) (result i32)
func (p *Pawtect) __wbindgen_is_object(args []wasmer.Value) ([]wasmer.Value, error) {
	log.Printf("__wbindgen_is_object({})", args)
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

// __wbg_process_dc0fbacc7c1c06f7 - (param externref) (result externref)
func (p *Pawtect) __wbg_process_dc0fbacc7c1c06f7(args []wasmer.Value) ([]wasmer.Value, error) {
	log.Printf("__wbg_process_dc0fbacc7c1c06f7({})", args)
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

// __wbg_versions_c01dfd4722a88165 - (param externref) (result externref)
func (p *Pawtect) __wbg_versions_c01dfd4722a88165(args []wasmer.Value) ([]wasmer.Value, error) {
	log.Printf("__wbg_versions_c01dfd4722a88165({})", args)
	return []wasmer.Value{wasmer.NewValue(nil, wasmer.AnyRef)}, nil
}

// __wbg_node_905d3e251edff8a2 - (param externref) (result externref)
func (p *Pawtect) __wbg_node_905d3e251edff8a2(arg0 uint32) uint32 {
	// Stub implementation for node object
	return arg0
}

// __wbindgen_is_string - (param externref) (result i32)
func (p *Pawtect) __wbindgen_is_string(arg0 uint32) int32 {
	// Stub implementation - return 1 if it's a string, 0 otherwise
	// For now, assume all externrefs are strings
	return 1
}

// __wbg_require_60cc747a6bc5215a - (result externref)
func (p *Pawtect) __wbg_require_60cc747a6bc5215a() uint32 {
	// Stub implementation for require function
	return 0
}

// __wbindgen_is_function - (param externref) (result i32)
func (p *Pawtect) __wbindgen_is_function(arg0 uint32) int32 {
	// Stub implementation - return 1 if it's a function, 0 otherwise
	// For now, assume all externrefs are functions
	return 1
}

// __wbindgen_string_new - (param i32 i32) (result externref)
func (p *Pawtect) __wbindgen_string_new(arg0, arg1 uint32) uint32 {
	// Stub implementation for creating new string
	return 0
}

// __wbg_msCrypto_a61aeb35a24c1329 - (param externref) (result externref)
func (p *Pawtect) __wbg_msCrypto_a61aeb35a24c1329(arg0 uint32) uint32 {
	// Stub implementation for msCrypto object
	return arg0
}

// __wbg_randomFillSync_ac0988aba3254290 - (param externref externref)
func (p *Pawtect) __wbg_randomFillSync_ac0988aba3254290(arg0, arg1 uint32) {
	// Stub implementation for randomFillSync
}

// __wbg_getRandomValues_b8f5dbd5f3995a9e - (param externref externref)
func (p *Pawtect) __wbg_getRandomValues_b8f5dbd5f3995a9e(arg0, arg1 uint32) {
	// Stub implementation for getRandomValues
}

// __wbg_newnoargs_105ed471475aaf50 - (param i32 i32) (result externref)
func (p *Pawtect) __wbg_newnoargs_105ed471475aaf50(arg0, arg1 uint32) uint32 {
	// Stub implementation for new constructor with no args
	return 0
}

// __wbg_call_672a4d21634d4a24 - (param externref externref) (result externref)
func (p *Pawtect) __wbg_call_672a4d21634d4a24(arg0, arg1 uint32) uint32 {
	// Stub implementation for function call with 2 args
	return 0
}

// __wbindgen_is_undefined - (param externref) (result i32)
func (p *Pawtect) __wbindgen_is_undefined(arg0 uint32) int32 {
	// Stub implementation - return 1 if undefined, 0 otherwise
	// For now, assume nothing is undefined
	return 0
}

// __wbg_call_7cccdd69e0791ae2 - (param externref externref externref) (result externref)
func (p *Pawtect) __wbg_call_7cccdd69e0791ae2(arg0, arg1, arg2 uint32) uint32 {
	// Stub implementation for function call with 3 args
	return 0
}

// __wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0 - (result i32)
func (p *Pawtect) __wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0() uint32 {
	// Stub implementation for GLOBAL_THIS accessor
	return 0
}

// __wbg_static_accessor_SELF_37c5d418e4bf5819 - (result i32)
func (p *Pawtect) __wbg_static_accessor_SELF_37c5d418e4bf5819() uint32 {
	// Stub implementation for SELF accessor
	return 0
}

// __wbg_static_accessor_WINDOW_5de37043a91a9c40 - (result i32)
func (p *Pawtect) __wbg_static_accessor_WINDOW_5de37043a91a9c40() uint32 {
	// Stub implementation for WINDOW accessor
	return 0
}

// __wbg_static_accessor_GLOBAL_88a902d13a557d07 - (result i32)
func (p *Pawtect) __wbg_static_accessor_GLOBAL_88a902d13a557d07() uint32 {
	// Stub implementation for GLOBAL accessor
	return 0
}

// __wbg_buffer_609cc3eee51ed158 - (param externref) (result externref)
func (p *Pawtect) __wbg_buffer_609cc3eee51ed158(arg0 uint32) uint32 {
	// Stub implementation for buffer accessor
	return arg0
}

// __wbg_newwithbyteoffsetandlength_d97e637ebe145a9a - (param externref i32 i32) (result externref)
func (p *Pawtect) __wbg_newwithbyteoffsetandlength_d97e637ebe145a9a(arg0, arg1, arg2 uint32) uint32 {
	// Stub implementation for new with byte offset and length
	return 0
}

// __wbg_new_a12002a7f91c75be - (param externref) (result externref)
func (p *Pawtect) __wbg_new_a12002a7f91c75be(arg0 uint32) uint32 {
	// Stub implementation for new constructor
	return 0
}

// __wbg_set_65595bdd868b3009 - (param externref externref i32)
func (p *Pawtect) __wbg_set_65595bdd868b3009(arg0, arg1, arg2 uint32) {
	// Stub implementation for set operation
}

// __wbg_newwithlength_a381634e90c276d4 - (param i32) (result externref)
func (p *Pawtect) __wbg_newwithlength_a381634e90c276d4(arg0 uint32) uint32 {
	// Stub implementation for new with length
	return 0
}

// __wbg_subarray_aa9065fa9dc5df96 - (param externref i32 i32) (result externref)
func (p *Pawtect) __wbg_subarray_aa9065fa9dc5df96(arg0, arg1, arg2 uint32) uint32 {
	// Stub implementation for subarray
	return 0
}

// __wbindgen_throw - (param i32 i32)
func (p *Pawtect) __wbindgen_throw(arg0, arg1 uint32) {
	// Stub implementation for throw
}

// __wbindgen_memory - (result externref)
func (p *Pawtect) __wbindgen_memory() uint32 {
	// Stub implementation for memory accessor
	return 0
}

// __wbindgen_init_externref_table - (func)
func (p *Pawtect) __wbindgen_init_externref_table() {
	// Stub implementation for externref table initialization
}

func (p *Pawtect) registerWBG(importObj *wasmer.ImportObject) {
	importObj.Register("wbg", map[string]wasmer.IntoExtern{
		"__wbg_crypto_574e78ad8b13b65f": wasmer.NewFunction(
			p.store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			p.__wbg_crypto_574e78ad8b13b65f,
		),
		"__wbindgen_is_object": wasmer.NewFunction(
			p.store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.I32),
			),
			p.__wbindgen_is_object,
		),
		"__wbg_process_dc0fbacc7c1c06f7": wasmer.NewFunction(
			p.store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			p.__wbg_process_dc0fbacc7c1c06f7,
		),
		"__wbg_versions_c01dfd4722a88165": wasmer.NewFunction(
			p.store,
			wasmer.NewFunctionType(
				wasmer.NewValueTypes(wasmer.AnyRef),
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			p.__wbg_versions_c01dfd4722a88165,
		),
		"__wbg_node_905d3e251edff8a2": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_node_905d3e251edff8a2(uint32(args[0].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbindgen_is_string": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbindgen_is_string(uint32(args[0].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_require_60cc747a6bc5215a": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(),
				i32Slice(wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_require_60cc747a6bc5215a()
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbindgen_is_function": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbindgen_is_function(uint32(args[0].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbindgen_string_new": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32, wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbindgen_string_new(uint32(args[0].ToI32()), uint32(args[1].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_msCrypto_a61aeb35a24c1329": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_msCrypto_a61aeb35a24c1329(uint32(args[0].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_randomFillSync_ac0988aba3254290": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(),
				i32Slice(wasmer.I32, wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				p.__wbg_randomFillSync_ac0988aba3254290(uint32(args[0].ToI32()), uint32(args[1].ToI32()))
				return nil, nil
			},
		),
		"__wbg_getRandomValues_b8f5dbd5f3995a9e": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(),
				i32Slice(wasmer.I32, wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				p.__wbg_getRandomValues_b8f5dbd5f3995a9e(uint32(args[0].ToI32()), uint32(args[1].ToI32()))
				return nil, nil
			},
		),
		"__wbg_newnoargs_105ed471475aaf50": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32, wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_newnoargs_105ed471475aaf50(uint32(args[0].ToI32()), uint32(args[1].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_call_672a4d21634d4a24": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32, wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_call_672a4d21634d4a24(uint32(args[0].ToI32()), uint32(args[1].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbindgen_is_undefined": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbindgen_is_undefined(uint32(args[0].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_call_7cccdd69e0791ae2": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32, wasmer.I32, wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_call_7cccdd69e0791ae2(uint32(args[0].ToI32()), uint32(args[1].ToI32()), uint32(args[2].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_static_accessor_GLOBAL_THIS_56578be7e9f832b0()
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_static_accessor_SELF_37c5d418e4bf5819": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_static_accessor_SELF_37c5d418e4bf5819()
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_static_accessor_WINDOW_5de37043a91a9c40": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_static_accessor_WINDOW_5de37043a91a9c40()
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_static_accessor_GLOBAL_88a902d13a557d07": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_static_accessor_GLOBAL_88a902d13a557d07()
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_buffer_609cc3eee51ed158": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_buffer_609cc3eee51ed158(uint32(args[0].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_newwithbyteoffsetandlength_d97e637ebe145a9a": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32, wasmer.I32, wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_newwithbyteoffsetandlength_d97e637ebe145a9a(uint32(args[0].ToI32()), uint32(args[1].ToI32()), uint32(args[2].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_new_a12002a7f91c75be": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_new_a12002a7f91c75be(uint32(args[0].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_set_65595bdd868b3009": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(),
				i32Slice(wasmer.I32, wasmer.I32, wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				p.__wbg_set_65595bdd868b3009(uint32(args[0].ToI32()), uint32(args[1].ToI32()), uint32(args[2].ToI32()))
				return nil, nil
			},
		),
		"__wbg_newwithlength_a381634e90c276d4": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_newwithlength_a381634e90c276d4(uint32(args[0].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbg_subarray_aa9065fa9dc5df96": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(wasmer.I32, wasmer.I32, wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbg_subarray_aa9065fa9dc5df96(uint32(args[0].ToI32()), uint32(args[1].ToI32()), uint32(args[2].ToI32()))
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbindgen_throw": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(),
				i32Slice(wasmer.I32, wasmer.I32),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				p.__wbindgen_throw(uint32(args[0].ToI32()), uint32(args[1].ToI32()))
				return nil, nil
			},
		),
		"__wbindgen_memory": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(wasmer.I32),
				i32Slice(),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				result := p.__wbindgen_memory()
				return []wasmer.Value{wasmer.NewI32Value(int32(result))}, nil
			},
		),
		"__wbindgen_init_externref_table": wasmer.NewFunction(
			p.engine,
			wasmer.NewFunctionType(
				i32Slice(),
				i32Slice(),
			),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				p.__wbindgen_init_externref_table()
				return nil, nil
			},
		),
	})
}
