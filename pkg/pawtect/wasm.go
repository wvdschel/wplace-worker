package pawtect

import (
	_ "embed"
	"log"

	"github.com/wasmerio/wasmer-go/wasmer"
)

//go:embed pawtect.wasm
var pawtectWasmModule []byte

type Pawtect struct {
	engine   *wasmer.Engine
	instance *wasmer.Instance
	store    *wasmer.Store
}

func Load() (*Pawtect, error) {
	p := &Pawtect{}

	p.engine = wasmer.NewEngine()
	p.store = wasmer.NewStore(p.engine)

	module, err := wasmer.NewModule(p.store, pawtectWasmModule)
	if err != nil {
		return nil, err
	}

	importObject := wasmer.NewImportObject()
	p.registerWBG(importObject)

	p.instance, err = wasmer.NewInstance(module, importObject)
	if err != nil {
		return nil, err
	}

	startFunc, err := p.instance.Exports.GetFunction("__wbindgen_start")
	if err != nil {
		log.Printf("failed to get __wbindgen_start: %s", err.Error())
		return nil, err
	}

	_, err = startFunc()
	if err != nil {
		log.Printf("failed to call __wbindgen_start: %s", err.Error())
		return nil, err
	}

	return p, nil
}
