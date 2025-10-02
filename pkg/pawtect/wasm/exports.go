package wasm

import "github.com/wasmerio/wasmer-go/wasmer"

func (o *Wasm) Memory() *wasmer.Memory {
	mem, _ := o.Instance.Exports.GetMemory("memory")
	return mem
}

func (o *Wasm) SetUserId(args ...any) (any, error) {
	wasmFunc, err := o.Instance.Exports.GetFunction("set_user_id")
	if err != nil { return nil, err }
	return wasmFunc(args...)
}

func (o *Wasm) RequestUrl(args ...any) (any, error) {
	wasmFunc, err := o.Instance.Exports.GetFunction("request_url")
	if err != nil { return nil, err }
	return wasmFunc(args...)
}

func (o *Wasm) GetLoadPayload(args ...any) (any, error) {
	wasmFunc, err := o.Instance.Exports.GetFunction("get_load_payload")
	if err != nil { return nil, err }
	return wasmFunc(args...)
}

func (o *Wasm) GetPawtectedEndpointPayload(args ...any) (any, error) {
	wasmFunc, err := o.Instance.Exports.GetFunction("get_pawtected_endpoint_payload")
	if err != nil { return nil, err }
	return wasmFunc(args...)
}

func (o *Wasm) WbindgenExnStore(args ...any) (any, error) {
	wasmFunc, err := o.Instance.Exports.GetFunction("__wbindgen_exn_store")
	if err != nil { return nil, err }
	return wasmFunc(args...)
}

func (o *Wasm) ExternrefTableAlloc(args ...any) (any, error) {
	wasmFunc, err := o.Instance.Exports.GetFunction("__externref_table_alloc")
	if err != nil { return nil, err }
	return wasmFunc(args...)
}

func (o *Wasm) WbindgenExport2() *wasmer.Table {
	table, _ := o.Instance.Exports.GetTable("__wbindgen_export_2")
	return table
}

func (o *Wasm) WbindgenMalloc(args ...any) (any, error) {
	wasmFunc, err := o.Instance.Exports.GetFunction("__wbindgen_malloc")
	if err != nil { return nil, err }
	return wasmFunc(args...)
}

func (o *Wasm) WbindgenRealloc(args ...any) (any, error) {
	wasmFunc, err := o.Instance.Exports.GetFunction("__wbindgen_realloc")
	if err != nil { return nil, err }
	return wasmFunc(args...)
}

func (o *Wasm) WbindgenFree(args ...any) (any, error) {
	wasmFunc, err := o.Instance.Exports.GetFunction("__wbindgen_free")
	if err != nil { return nil, err }
	return wasmFunc(args...)
}

func (o *Wasm) WbindgenStart(args ...any) (any, error) {
	wasmFunc, err := o.Instance.Exports.GetFunction("__wbindgen_start")
	if err != nil { return nil, err }
	return wasmFunc(args...)
}

