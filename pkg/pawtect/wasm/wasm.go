package wasm

import "github.com/wasmerio/wasmer-go/wasmer"

type Wasm struct {
	Store *wasmer.Store
	Instance *wasmer.Instance
	Module *wasmer.Module
}

func New(store *wasmer.Store, module *wasmer.Module) (*Wasm, error) {
	res := &Wasm{
		Store: store,
		Module: module,
	}

	importObject := wasmer.NewImportObject()
	res.Registerwbg(importObject)

	inst, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		return nil, err
	}

	res.Instance = inst

	return res, nil
}
